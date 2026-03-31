package domain

import (
	"time"

	"github.com/4strodev/4stroblog/site/shared/db/models"
	"github.com/google/uuid"
)

type SessionBuilder struct {
	expirationTime time.Time
}

// ExpirateAfter sets the expirationTime to now() + duration
func (b *SessionBuilder) ExpirateAfter(duration time.Duration) *SessionBuilder {
	b.expirationTime = time.Now().Add(duration)
	return b
}

// SetExpirationTime sets the expiration time to the provided time
func (b *SessionBuilder) SetExpirationTime(time time.Time) *SessionBuilder {
	b.expirationTime = time
	return b
}

// Build returns a new session. At least profile must be set before calling
// build. If no expirationTime is set the deafult one is 5 hours
func (b *SessionBuilder) Build(profile models.Profile) (Session, error) {
	var session Session

	if b.expirationTime.IsZero() {
		b.ExpirateAfter(time.Hour * 5)
	}

	// Create refresh token

	session = Session{
		ID:             uuid.Must(uuid.NewV7()),
		UserID:         profile.UserID,
		ProfileID:      profile.ID,
		ExpriationTime: b.expirationTime,
	}

	return session, nil
}
