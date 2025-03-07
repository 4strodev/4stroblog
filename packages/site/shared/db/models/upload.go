package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Upload struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primaryKey"`
	Hash     string
	Name     string
	MimeType string
	Time     time.Time
}
