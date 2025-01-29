package services

import (
	"errors"
	"fmt"

	"github.com/4strodev/4stroblog/site/shared/db/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterService struct {
	DB *gorm.DB
}

type RegisterReqDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserRegisterResDTO struct {
	UserID uuid.UUID `json:"userId"`
}

func (s *RegisterService) Register(req RegisterReqDTO) (res UserRegisterResDTO, err error) {
	profile := models.Profile{}
	var count int64
	err = s.DB.Model(&profile).Where("email = ?", req.Email).Limit(1).Count(&count).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return res, fmt.Errorf("error fetching profiels from db: %w", err)
		}
	}
	if count > 0 {
		return res, fmt.Errorf("profile with email '%s' already exists", req.Email)
	}

	// Create user
	err = s.DB.Transaction(func(tx *gorm.DB) error {
		// Create user
		user := models.User{
			ID: uuid.Must(uuid.NewV7()),
		}
		err = s.DB.Create(&user).Error
		if err != nil {
			return fmt.Errorf("error creating user: %w", err)
		}
		res.UserID = user.ID

		// Create profile
		profile.Email = req.Email
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
		if err != nil {
			return fmt.Errorf("error hashing password: %w", err)
		}
		profile.Password = string(hash)
		profile.ID = uuid.Must(uuid.NewV7())
		profile.Name = req.Name
		profile.User = user
		err = s.DB.Create(&profile).Error
		if err != nil {
			return fmt.Errorf("cannot create a user profile: %w", err)
		}
		return nil
	})

	return
}
