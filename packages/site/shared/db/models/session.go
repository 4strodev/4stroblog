package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID             uuid.UUID `gorm:"primaryKey;column:id"`
	UserID         uuid.UUID `gorm:"column:user_id"`
	User           User
	ProfileID      uuid.UUID `gorm:"column:profile_id"`
	Profile        Profile
	ExpirationTime time.Time `gorm:"column:expiration_time"`
}
