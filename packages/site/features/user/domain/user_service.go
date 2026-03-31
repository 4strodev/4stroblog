package domain

import (
	"context"

	domainerrors "github.com/4strodev/4stroblog/site/shared/domain/domainerrors"
)

type UserService struct {
	Repository UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, user User) error {
	user, err := s.Repository.FindByEmail(ctx, user.PrimaryEmail)
	if err != nil {
		_, isNotFound := domainerrors.Is(err, domainerrors.ENTITY_NOT_FOUND)
		if !isNotFound {
			return err
		}
	} else {
		return domainerrors.Errorf(
			domainerrors.DATA_CONFLICT,
			"User with email '%s' already exists",
			user.PrimaryEmail)
	}
	return s.Repository.Save(ctx, user)
}

// Verify user implies to send an email verification and ensure that user
// clicked the verification link
func (s *UserService) VerifyUser(ctx context.Context, user User) error {
	user.Verified = true
	return s.Repository.Save(ctx, user)
}
