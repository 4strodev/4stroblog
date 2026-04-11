package domain

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	// Save persists the given user. Returns a DATABASE domain error on failure.
	Save(ctx context.Context, user User) error
	// FindByID queries for a non-deleted user by UUID.
	// Returns ENTITY_NOT_FOUND if no user exists with that ID, DATABASE on other failures.
	FindByID(ctx context.Context, id uuid.UUID) (User, error)
}
