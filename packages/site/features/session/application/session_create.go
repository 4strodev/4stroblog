package application

import (
	"context"

	"github.com/4strodev/4stroblog/site/features/session/application/dto"
	"github.com/4strodev/4stroblog/site/features/session/domain"
	"github.com/4strodev/4stroblog/site/shared/domain/domainerrors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type SessionCreateReq struct {
	ID       uuid.UUID `json:"id"`
	User     string    `json:"user"`
	Password string    `json:"password"`
}

func (s *SessionAppService) Create(ctx context.Context, req SessionCreateReq) (dto.SessionDto, error) {
	var sessionDto dto.SessionDto

	// Look up the profile by email to get the associated user ID.
	profile, err := s.ProfileRepository.FindByEmail(ctx, req.User)
	if err != nil {
		return sessionDto, err
	}

	// Fetch the user record to retrieve the password hash.
	user, err := s.UserRepository.FindByID(ctx, profile.UserID)
	if err != nil {
		return sessionDto, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return sessionDto, domainerrors.WrapError(domainerrors.DATA_CONFLICT, err)
	}

	sessionBuilder := domain.SessionBuilder{}
	session, err := sessionBuilder.Build(profile)
	if err != nil {
		return sessionDto, err
	}

	if err := s.SessionService.Save(ctx, session); err != nil {
		return sessionDto, err
	}

	sessionDto = dto.SessionDto{
		ID:             session.ID,
		UserID:         session.UserID,
		ExpirationTime: session.ExpriationTime,
		ProfileID:      session.ProfileID,
	}
	return sessionDto, nil
}
