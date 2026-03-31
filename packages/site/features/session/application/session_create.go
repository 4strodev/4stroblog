package application

import (
	"context"
	"fmt"

	"github.com/4strodev/4stroblog/site/features/session/domain"
	"github.com/4strodev/4stroblog/site/shared/db/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SessionCreateReq struct {
	ID       uuid.UUID `json:"id"`
	User     string    `json:"user"`
	Password string    `json:"password"`
}

func (s *SessionAppService) Create(ctx context.Context, req SessionCreateReq) (err error) {
	email, password := req.User, req.Password
	var session domain.Session
	var profile models.Profile

	// Getting user profile
	session, err = s.SessionService.FindByEmail(ctx, email)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte(password))
	if err != nil {
		err = fmt.Errorf("password does not match: %w", err)
		return
	}

	sessionBuilder := domain.SessionBuilder{}
	session, err = sessionBuilder.Build(profile)
	if err != nil {
		return
	}

	// Saving session to database
	return s.SessionService.Save(ctx, session)
}
