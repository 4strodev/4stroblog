package session

import (
	"time"

	"github.com/google/uuid"
)

// Session is the entity that handles logic related to [Session].
// A session stores information about a user that has logged in
// to the application. Stores their duration, user id.
type Session struct {
	ID             uuid.UUID
	UserID         uuid.UUID
	ExpriationTime time.Time
	ProfileID      uuid.UUID
}

func (s *Session) HasExpired() bool {
	return s.ExpriationTime.After(time.Now())
}
