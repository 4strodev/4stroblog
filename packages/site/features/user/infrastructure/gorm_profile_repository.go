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

type GormProfileRepository struct {
	DB *gorm.DB
}

func (r *GormProfileRepository) Save(ctx context.Context, profile domain.Profile) error {
	profileModel := models.Profile{
		ID:     profile.ID,
		UserID: profile.UserID,
		Email:  profile.Email,
		Name:   profile.Name,
	}

	err := r.DB.WithContext(ctx).Create(&profileModel).Error
	if err != nil {
		return domainerrors.WrapError(domainerrors.DATABASE, err)
	}
	return nil
}

// FindByEmail uses First which returns gorm.ErrRecordNotFound when no row matches.
// That is mapped to ENTITY_NOT_FOUND so callers can distinguish "not found" from other DB errors.
func (r *GormProfileRepository) FindByEmail(ctx context.Context, email string) (domain.Profile, error) {
	profileModel := models.Profile{}
	profile := domain.Profile{}

	err := r.DB.WithContext(ctx).
		Where("email = ?", email).
		First(&profileModel).Error
	if err != nil {
		errorCode := domainerrors.DATABASE
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorCode = domainerrors.ENTITY_NOT_FOUND
		}
		return profile, domainerrors.WrapError(errorCode, err)
	}

	profile = domain.Profile{
		ID:     profileModel.ID,
		UserID: profileModel.UserID,
		Email:  profileModel.Email,
		Name:   profileModel.Name,
	}
	return profile, nil
}

func (r *GormProfileRepository) FindByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Profile, error) {
	var profileModels []models.Profile

	err := r.DB.WithContext(ctx).
		Where("user_id = ? AND deleted_at IS NULL", userID).
		Find(&profileModels).Error
	if err != nil {
		errorCode := domainerrors.DATABASE
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorCode = domainerrors.ENTITY_NOT_FOUND
		}
		return nil, domainerrors.WrapError(errorCode, err)
	}

	profiles := make([]domain.Profile, 0, len(profileModels))
	for _, m := range profileModels {
		profiles = append(profiles, domain.Profile{
			ID:     m.ID,
			UserID: m.UserID,
			Email:  m.Email,
			Name:   m.Name,
		})
	}
	return profiles, nil
}
