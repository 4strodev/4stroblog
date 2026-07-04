package infrastructure

import (
	"context"
	"github.com/4strodev/4stroblog/site/features/user/domain"
	"github.com/google/uuid"
)

type InMemoryUserRepository struct {
	InnerMap map[uuid.UUID]domain.User
}

// FindByID implements [domain.UserRepository].
func (i *InMemoryUserRepository) FindByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	panic("unimplemented")
}

// Save implements [domain.UserRepository].
func (i *InMemoryUserRepository) Save(ctx context.Context, user domain.User) error {
	panic("unimplemented")
}
