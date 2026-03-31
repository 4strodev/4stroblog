package domain

import "context"

type UserRepository interface {
	// Save persists the given user. Returns a DATABASE domain error on failure.
	Save(ctx context.Context, user User) error
	// FindByEmail queries for a non-deleted user whose email matches the given value.
	// Returns a User and nil error on success.
	// Returns a domain error on failure
	FindByEmail(ctx context.Context, email string) (User, error)
}
