package domain

import (
	"context"

	"github.com/google/uuid"
)

type ProfileRepository interface {
	// Save persists the given profile. Returns a DATABASE domain error on failure.
	Save(ctx context.Context, profile Profile) error
	// FindByEmail queries for a non-deleted profile whose email matches the given value.
	// Returns ENTITY_NOT_FOUND if no profile exists with that email, DATABASE on other failures.
	FindByEmail(ctx context.Context, email string) (Profile, error)
	// FindByUserID returns all non-deleted profiles associated with the given user UUID.
	FindByUserID(ctx context.Context, userID uuid.UUID) ([]Profile, error)
}
