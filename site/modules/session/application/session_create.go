package application

import (
	"errors"
	"fmt"
	"time"

	appSession "github.com/4strodev/4stroblog/site/modules/session/domain"
	"github.com/4strodev/4stroblog/site/shared/db/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type SessionCreateReq struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type SessionCreateRes struct {
	ExpirationTime time.Time `json:"expirationTime"`
	ID             uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"userId"`
	ProfileID      uuid.UUID `json:"profileId"`
}

func (s *SessionService) Create(req SessionCreateReq) (SessionCreateRes, error) {
	email, password := req.User, req.Password
	var session = appSession.Session{}
	var response SessionCreateRes
	var profile = models.Profile{}

	// Getting user profile
	err := s.DB.First(&profile, "email = ?", email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, fmt.Errorf("user not found")
		}
		return response, fmt.Errorf("error getting user from database: %w", err)
	}

	// Checking password match
	err = bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte(password))
	if err != nil {
		return response, fmt.Errorf("password does not match: %w", err)
	}

	// Building session
	sessionBuilder := appSession.SessionBuilder{}
	session, err = sessionBuilder.Build(profile)
	if err != nil {
		return response, err
	}
	sessionModel := models.Session{
		ID:             session.ID,
		UserID:         session.UserID,
		ExpirationTime: session.ExpriationTime,
	}

	s.DB.Create(&sessionModel)
	response = SessionCreateRes{
		ID:             session.ID,
		UserID:         session.UserID,
		ExpirationTime: session.ExpriationTime,
		ProfileID:      session.ProfileID,
	}
	return response, nil
}
