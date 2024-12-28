package services

import (
	"errors"
	"fmt"
	"time"

	appSession "github.com/4strodev/4stroblog/site/application/session"
	"github.com/4strodev/4stroblog/site/shared/config"
	"github.com/4strodev/4stroblog/site/shared/db/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginService struct {
	DB     *gorm.DB
	Config config.Config
}

type LoginReqDTO struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type LoginSessionResDTO struct {
	ExpirationTime time.Time `json:"expirationTime"`
	ID             uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"userId"`
	ProfileID      uuid.UUID `json:"profileId"`
}

type LoginResDTO struct {
	Session      LoginSessionResDTO `json:"session"`
	AccessToken  string             `json:"accessToken"`
	RefreshToken string             `json:"refreshToken"`
}

func (s *LoginService) Login(req LoginReqDTO) (LoginResDTO, error) {
	email, password := req.User, req.Password
	response := LoginResDTO{}
	session := appSession.Session{}
	profile := models.Profile{}

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
	// Creating JWTs
	jwtBuilder := appSession.JWTBuilder{}
	jwts, err := jwtBuilder.SetSecret(s.Config.JWK.Secret).Build(session)
	if err != nil {
		return response, err
	}

	sessionModel := models.Session{
		ID:             session.ID,
		UserID:         session.UserID,
		ExpirationTime: session.ExpriationTime,
	}

	s.DB.Create(&sessionModel)
	response = LoginResDTO{
		Session: LoginSessionResDTO{
			ID:             session.ID,
			UserID:         session.UserID,
			ExpirationTime: session.ExpriationTime,
			ProfileID:      session.ProfileID,
		},
		AccessToken:  jwts.AccessToken,
		RefreshToken: jwts.RefreshToken,
	}

	return response, nil
}
