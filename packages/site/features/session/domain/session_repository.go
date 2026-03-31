package domain

import (
	"context"

	"github.com/google/uuid"
)

type SessionRepository interface {
	Save(ctx context.Context, session Session) error
	FindById(ctx context.Context, id uuid.UUID) (Session, error)
	FindByProfileId(ctx context.Context, userId uuid.UUID) (Session, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
