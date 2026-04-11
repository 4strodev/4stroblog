package domain

import (
	"time"

	"github.com/4strodev/4stroblog/site/shared/domain/domainerrors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID
	Password  string // bcrypt hash
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewUser(password string) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return User{}, domainerrors.WrapError(domainerrors.RUNTIME, err)
	}
	return User{
		ID:       uuid.Must(uuid.NewV7()),
		Password: string(hash),
	}, nil
}
