package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Upload struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primaryKey"`
	Hash     sql.NullString
	Name     string
	MimeType string
	Time     time.Time
}
