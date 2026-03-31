package domain

import (
	"context"

	"github.com/google/uuid"
)

type SessionRepository interface {
	// Save persists the given session, creating it if it does not exist or updating it if it does.
	// Returns a DATABASE domain error on failure.
	Save(ctx context.Context, session Session) error
	// FindById queries for a session by its UUID.
	// Returns ENTITY_NOT_FOUND if no session exists with that ID, DATABASE on other failures.
	FindById(ctx context.Context, id uuid.UUID) (Session, error)
	// FindByProfileId queries for the session associated with the given profile UUID.
	// Returns ENTITY_NOT_FOUND if no session exists for that profile, DATABASE on other failures.
	FindByProfileId(ctx context.Context, profileId uuid.UUID) (Session, error)
	// Delete removes the session with the given UUID. Returns a DATABASE domain error on failure.
	Delete(ctx context.Context, id uuid.UUID) error
}
