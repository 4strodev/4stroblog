package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID             uuid.UUID `gorm:"primaryKey"`
	UserID         uuid.UUID
	User           User
	ProfileID      uuid.UUID
	Profile        Profile
	ExpirationTime time.Time
}
