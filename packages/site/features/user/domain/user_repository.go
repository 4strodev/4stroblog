package domain

import "context"

type UserRepository interface {
	Save(ctx context.Context, user User) error
	// FindByEmail looks for an alive user with this primary email
	// if user is not found then an error is returned
	FindByEmail(ctx context.Context, email string) (User, error)
}
