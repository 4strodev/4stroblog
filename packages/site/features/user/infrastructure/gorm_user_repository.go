package infrastructure

import (
	"context"
	"errors"

	"github.com/4strodev/4stroblog/site/features/user/domain"
	"github.com/4strodev/4stroblog/site/shared/db/models"
	"github.com/4strodev/4stroblog/site/shared/domain/domainerrors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func (r *GormUserRepository) Save(ctx context.Context, user domain.User) error {
	userModel := models.User{
		ID:       user.ID,
		Password: user.Password,
	}

	err := r.DB.WithContext(ctx).Create(&userModel).Error
	if err != nil {
		return domainerrors.WrapError(domainerrors.DATABASE, err)
	}
	return nil
}

// FindByID uses First which returns gorm.ErrRecordNotFound when no row matches.
// That is mapped to ENTITY_NOT_FOUND so callers can distinguish "not found" from other DB errors.
func (r *GormUserRepository) FindByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	userModel := models.User{}
	user := domain.User{}

	err := r.DB.WithContext(ctx).First(&userModel, id).Error
	if err != nil {
		errorCode := domainerrors.DATABASE
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorCode = domainerrors.ENTITY_NOT_FOUND
		}
		return user, domainerrors.WrapError(errorCode, err)
	}

	user = domain.User{
		ID:       userModel.ID,
		Password: userModel.Password,
	}
	return user, nil
}
