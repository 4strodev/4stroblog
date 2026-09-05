package application

import (
	"context"
	"io"
	"time"

	"github.com/4strodev/4stroblog/site/features/uploads/domain"
	"github.com/4strodev/4stroblog/site/shared/config"
	"github.com/4strodev/4stroblog/site/shared/domain/domainerrors"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

func NewUploadsService(
	db *gorm.DB,
	s3 *minio.Client,
	uploadsRepository domain.UploadsRepository,
	config config.Config,
) *UploadsService {
	return &UploadsService{
		Db:                db,
		ObjectStorage:     s3,
		UploadsRepository: uploadsRepository,
		Config:            config,
	}
}

type UploadsService struct {
	Db                *gorm.DB
	ObjectStorage     *minio.Client
	UploadsRepository domain.UploadsRepository
	Config            config.Config
}

func (s *UploadsService) DeleteUpload(ctx context.Context, id uuid.UUID) error {
	upload, err := s.UploadsRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	err = s.UploadsRepository.DeleteById(ctx, id)
	if err != nil {
		return err
	}

	_, err = s.UploadsRepository.FindByHash(ctx, upload.Hash)
	domainErr, isNotFound := domainerrors.Is(err, domainerrors.ENTITY_NOT_FOUND)
	if !isNotFound && err != nil {
		return domainErr
	}

	if isNotFound {
		return nil
	}

	// No more uploads with this hash remove object from storage
	err = s.ObjectStorage.RemoveObject(
		ctx,
		s.Config.Storage.S3.Bucket,
		upload.StringHash(),
		minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

// SaveUpload saves metadata and blob. If blob is already persisted it omits it's persist step.
// it updates inner upload data that is not set: uuid, name, time, etc.
// it can return a [domainerrors.RUNTIME]
func (s *UploadsService) SaveUpload(ctx context.Context, upload *domain.Upload, blob io.ReadSeeker) error {
	err := upload.DigestBlob(blob)
	if err != nil {
		return domainerrors.Errorf(domainerrors.RUNTIME, "cannot calculate upload hash: %w", err)
	}
	_, err = blob.Seek(0, io.SeekStart)
	if err != nil {
		return domainerrors.Errorf(domainerrors.RUNTIME, "cannot seek upload blob: %w", err)
	}

	if upload.ID == uuid.Nil {
		upload.ID = uuid.New()
	}

	if upload.Name == "" {
		upload.Name = upload.StringHash()
	}

	if upload.Time.IsZero() {
		upload.Time = time.Now()
	}

	_, err = s.UploadsRepository.FindByHash(ctx, upload.Hash)
	domainErr, isNotFound := domainerrors.Is(err, domainerrors.ENTITY_NOT_FOUND)
	if !isNotFound && err != nil {
		return domainErr
	}

	if isNotFound {
		_, err = s.ObjectStorage.PutObject(ctx,
			s.Config.Storage.S3.Bucket,
			upload.StringHash(),
			blob,
			int64(upload.Size),
			minio.PutObjectOptions{
				ContentType: upload.MimeType,
			})
		if err != nil {
			return domainerrors.Errorf(domainerrors.STORAGE, "cannot put upload blob: %w", err)
		}
	}

	err = s.UploadsRepository.Save(ctx, *upload)
	if err != nil {
		return domainerrors.Errorf(domainerrors.DATABASE, "cannot save upload meta-data: %w", err)
	}

	return nil
}
