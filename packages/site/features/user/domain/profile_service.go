package domain

import (
	"context"

	"github.com/4strodev/4stroblog/site/shared/domain/domainerrors"
	"github.com/google/uuid"
)

type ProfileService struct {
	ProfileRepository ProfileRepository
}

// CreateProfile creates a new profile for an existing user.
// It returns DATA_CONFLICT if a profile with the given email already exists.
func (s *ProfileService) CreateProfile(ctx context.Context, userID uuid.UUID, name, email string) (Profile, error) {
	var profile Profile

	_, err := s.ProfileRepository.FindByEmail(ctx, email)
	_, notFound := domainerrors.Is(err, domainerrors.ENTITY_NOT_FOUND)
	if err != nil && !notFound {
		return profile, err
	}
	if !notFound {
		return profile, domainerrors.Errorf(domainerrors.DATA_CONFLICT, "a profile with email %q already exists", email)
	}

	profile = NewProfile(userID, email, name)
	if err := s.ProfileRepository.Save(ctx, profile); err != nil {
		return profile, err
	}

	return profile, nil
}
