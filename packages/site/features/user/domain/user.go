package domain

import (
	"fmt"
	"slices"

	"github.com/4strodev/4stroblog/site/shared/domain/domainerrors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID   uuid.UUID
	Name string
	// Login it's the identifier that will be used to login
	Login        string
	PrimaryEmail string
	Password     string
	Verified     bool
	Emails       []string
}

func NewUser(name string, login string, password string) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return User{}, domainerrors.WrapError(domainerrors.RUNTIME, err)
	}
	return User{
		ID:       uuid.Must(uuid.NewV7()),
		Name:     name,
		Login:    login,
		Password: string(passwordHash),
	}, nil
}

// TODO AddEmail
// TODO RemoveEmail

func (u *User) SetPrimaryEmail(email string) error {
	if slices.Contains(u.Emails, email) {
		u.PrimaryEmail = email
		return nil
	}

	return fmt.Errorf("cannot set primary email '%s': email not saved", email)
}
