package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	ID       uuid.UUID
	UserID   uuid.UUID
	User     User
	Email    string
	Password string
	Name     string
}
