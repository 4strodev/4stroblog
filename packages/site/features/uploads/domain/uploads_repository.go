package domain

import (
	"context"

	"github.com/google/uuid"
)

type UploadsRepository interface {
	// Save persists the given upload metadata. Returns a DATABASE domain error on failure.
	Save(ctx context.Context, upload Upload) error
	// FindByID queries for an upload by UUID.
	// Returns ENTITY_NOT_FOUND if no upload exists with that ID, DATABASE on other failures.
	FindByID(ctx context.Context, id uuid.UUID) (Upload, error)
}
