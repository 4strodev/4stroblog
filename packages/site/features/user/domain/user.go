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
	// PrimaryEmail is the email used for login
	// is the human readable identifier for this profile
	PrimaryEmail string
	Password     string
	Verified     bool
	Emails       []string
}

func NewUser(name string, primaryEmail string, password string) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return User{}, domainerrors.WrapError(domainerrors.RUNTIME, err)
	}
	return User{
		ID:           uuid.Must(uuid.NewV7()),
		Name:         name,
		PrimaryEmail: primaryEmail,
		Password:     string(passwordHash),
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
