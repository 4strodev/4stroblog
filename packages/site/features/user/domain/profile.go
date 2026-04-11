package domain

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Email     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewProfile(userID uuid.UUID, email, name string) Profile {
	return Profile{
		ID:     uuid.Must(uuid.NewV7()),
		UserID: userID,
		Email:  email,
		Name:   name,
	}
}
