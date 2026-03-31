package dto

import (
	"time"

	"github.com/google/uuid"
)

type SessionDto struct {
	ID             uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"user_id"`
	ExpirationTime time.Time `json:"expiration_time"`
	ProfileID      uuid.UUID `json:"profile_id"`
}
