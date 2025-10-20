package infrastructure

import (
	"context"
	"errors"

	"github.com/4strodev/4stroblog/site/features/user/domain"
	"github.com/4strodev/4stroblog/site/shared/db/models"
	domainerrors "github.com/4strodev/4stroblog/site/shared/domain/domainerrors"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func (r *GormUserRepository) Save(ctx context.Context, user domain.User) error {
	userModel := models.User{
		ID:    user.ID,
		Email: user.PrimaryEmail,
	}

	err := r.DB.WithContext(ctx).Create(&userModel).Error
	if err != nil {
		return domainerrors.WrapError(domainerrors.DATABASE, err)
	}
	// TODO create profile
	return nil
}
func (r *GormUserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	userModel := models.User{}
	user := domain.User{}
	err := r.DB.WithContext(ctx).Model(&models.User{}).
		Where("email = ? AND deleted_at IS NULL", email).
		First(&userModel).Error

	if err != nil {
		errorCode := domainerrors.DATABASE
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			errorCode = domainerrors.ENTITY_NOT_FOUND
		}
		return user, domainerrors.WrapError(errorCode, err)
	}

	user = domain.User{
		ID:           userModel.ID,
		PrimaryEmail: userModel.Email,
		Password:     userModel.Password,
		Verified:     userModel.Verified,
		Name:         userModel.Name,
		Emails:       make([]string, 0),
	}

	return user, nil
}
