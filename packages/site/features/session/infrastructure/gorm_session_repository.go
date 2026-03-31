package infrastructure

import (
	"context"
	"errors"

	"github.com/4strodev/4stroblog/site/features/session/domain"
	"github.com/4strodev/4stroblog/site/shared/db/models"
	"github.com/4strodev/4stroblog/site/shared/domain/domainerrors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormSessionRepository struct {
	DB *gorm.DB
}

// FindById implements domain.SessionRepository.
func (r *GormSessionRepository) FindById(ctx context.Context, id uuid.UUID) (domain.Session, error) {
	sessionModel := models.Session{}
	session := domain.Session{}

	err := r.DB.WithContext(ctx).First(&sessionModel, id).Error
	if err != nil {
		errorCode := domainerrors.DATABASE
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorCode = domainerrors.ENTITY_NOT_FOUND
		}
		return session, domainerrors.WrapError(errorCode, err)
	}

	session = sessionModelToSession(sessionModel)

	return session, nil
}

func (r *GormSessionRepository) FindByProfileId(ctx context.Context, profileId uuid.UUID) (domain.Session, error) {
	sessionModel := models.Session{}
	session := domain.Session{}

	err := r.DB.WithContext(ctx).First(&sessionModel, "profile_id = ?", profileId).Error
	if err != nil {
		errorCode := domainerrors.DATABASE
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorCode = domainerrors.ENTITY_NOT_FOUND
		}
		return session, domainerrors.WrapError(errorCode, err)
	}

	session = sessionModelToSession(sessionModel)

	return session, nil
}

func (r *GormSessionRepository) FindByEmail(ctx context.Context, email string) (domain.Session, error) {
	sessionModel := models.Session{}
	session := domain.Session{}

	err := r.DB.WithContext(ctx).First(&sessionModel, "email = ?", email).Error
	if err != nil {
		errorCode := domainerrors.DATABASE
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorCode = domainerrors.ENTITY_NOT_FOUND
		}
		return session, domainerrors.WrapError(errorCode, err)
	}

	session = sessionModelToSession(sessionModel)

	return session, nil
}

// Save implements domain.SessionRepository.
func (r *GormSessionRepository) Save(ctx context.Context, session domain.Session) error {
	sessionModel := models.Session{
		ID:             session.ID,
		ProfileID:      session.ProfileID,
		UserID:         session.UserID,
		ExpirationTime: session.ExpriationTime,
	}

	err := r.DB.WithContext(ctx).Save(&sessionModel).Error
	if err != nil {
		return domainerrors.WrapError(domainerrors.DATABASE, err)
	}

	return nil
}

func (r *GormSessionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.DB.WithContext(ctx).Delete(&models.Session{}, id).Error
	if err != nil {
		return domainerrors.WrapError(domainerrors.DATABASE, err)
	}

	return nil
}
