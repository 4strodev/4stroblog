package infrastructure

import (
	"context"
	"errors"

	"github.com/4strodev/4stroblog/site/features/uploads/domain"
	"github.com/4strodev/4stroblog/site/shared/db/models"
	"github.com/4strodev/4stroblog/site/shared/domain/domainerrors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormUploadsRepository struct {
	DB *gorm.DB
}

// FindByHash implements [domain.UploadsRepository].
func (r *GormUploadsRepository) FindByHash(ctx context.Context, hash []byte) (domain.Upload, error) {
	uploadModel := models.Upload{}
	upload := domain.Upload{}

	err := r.DB.WithContext(ctx).First(&uploadModel, "hash = ?", hash).Error
	if err != nil {
		errorCode := domainerrors.DATABASE
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorCode = domainerrors.ENTITY_NOT_FOUND
		}
		return upload, domainerrors.WrapError(errorCode, err)
	}

	upload = domain.Upload{
		ID:       uploadModel.ID,
		Hash:     uploadModel.Hash,
		Name:     uploadModel.Name,
		MimeType: uploadModel.MimeType,
		Time:     uploadModel.Time,
	}
	return upload, nil
}

func (r *GormUploadsRepository) Save(ctx context.Context, upload domain.Upload) error {
	uploadModel := models.Upload{
		ID:       upload.ID,
		Hash:     upload.Hash,
		Name:     upload.Name,
		MimeType: upload.MimeType,
		Time:     upload.Time,
	}

	err := r.DB.WithContext(ctx).Save(&uploadModel).Error
	if err != nil {
		return domainerrors.WrapError(domainerrors.DATABASE, err)
	}
	return nil
}

// FindByID uses First which returns gorm.ErrRecordNotFound when no row matches.
// That is mapped to ENTITY_NOT_FOUND so callers can distinguish "not found" from other DB errors.
func (r *GormUploadsRepository) FindByID(ctx context.Context, id uuid.UUID) (domain.Upload, error) {
	uploadModel := models.Upload{}
	upload := domain.Upload{}

	err := r.DB.WithContext(ctx).First(&uploadModel, "id = ?", id).Error
	if err != nil {
		errorCode := domainerrors.DATABASE
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorCode = domainerrors.ENTITY_NOT_FOUND
		}
		return upload, domainerrors.WrapError(errorCode, err)
	}

	upload = domain.Upload{
		ID:       uploadModel.ID,
		Hash:     uploadModel.Hash,
		Name:     uploadModel.Name,
		MimeType: uploadModel.MimeType,
		Time:     uploadModel.Time,
	}
	return upload, nil
}
