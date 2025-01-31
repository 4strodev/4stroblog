package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	appSession "github.com/4strodev/4stroblog/site/features/session/domain"
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

func (s *SessionService) Create(ctx context.Context, req SessionCreateReq) (response SessionCreateRes, err error) {
	email, password := req.User, req.Password
	var session appSession.Session
	var profile models.Profile

	// Getting user profile
	err = s.DB.WithContext(ctx).First(&profile, "email = ?", email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, fmt.Errorf("user not found")
		}
		return
	}

	// Checking password match
	if err = ctx.Err(); err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte(password))
	if err != nil {
		err = fmt.Errorf("password does not match: %w", err)
		return
	}

	// Building session
	if err = ctx.Err(); err != nil {
		return
	}
	sessionBuilder := appSession.SessionBuilder{}
	session, err = sessionBuilder.Build(profile)
	if err != nil {
		return
	}
	sessionModel := models.Session{
		ID:             session.ID,
		UserID:         session.UserID,
		ExpirationTime: session.ExpriationTime,
	}

	// Saving session to database
	err = s.DB.WithContext(ctx).Create(&sessionModel).Error
	if err != nil {
		return
	}
	response = SessionCreateRes{
		ID:             session.ID,
		UserID:         session.UserID,
		ExpirationTime: session.ExpriationTime,
		ProfileID:      session.ProfileID,
	}
	return
}
