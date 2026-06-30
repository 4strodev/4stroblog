package application

import (
	"bytes"
	"context"
	"errors"
	"io"
	"time"

	"github.com/4strodev/4stroblog/site/features/uploads/domain"
	"github.com/4strodev/4stroblog/site/shared/db/models"
	"github.com/4strodev/4stroblog/site/shared/s3"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

func NewUploadsService(db *gorm.DB, s3 *minio.Client) *UploadsService {
	return &UploadsService{
		Db:            db,
		ObjectStorage: s3,
	}
}

type UploadsService struct {
	Db            *gorm.DB
	ObjectStorage *minio.Client
}

// Saves a file directly to s3 returning an error if something happens. It does not do any modification
// to the upload content. It just adds the time if it is missing.
// If the content is nil it returns an error
func (s *UploadsService) SaveFile(ctx context.Context, uploadFile domain.Upload) error {
	if uploadFile.Content == nil {
		return errors.New("cannot upload an file with no content")
	}

	if uploadFile.Time.IsZero() {
		uploadFile.Time = time.Now()
	}

	var buffer bytes.Buffer
	_, err := io.Copy(&buffer, uploadFile.Content)
	if err != nil {
		return err
	}

	uploadInfo, err := s.ObjectStorage.PutObject(
		ctx,
		s3.UPLOADS_BUCKET,
		uploadFile.ID.String(),
		uploadFile.Content,
		int64(buffer.Len()),
		minio.PutObjectOptions{Checksum: minio.ChecksumSHA256,
			UserMetadata: map[string]string{
				"name": uploadFile.Name,
				"time": uploadFile.Time.UTC().String(),
			}},
	)
	if err != nil {
		return err
	}
	uploadFile.Hash = uploadInfo.ChecksumSHA256

	uploadModel := models.Upload{
		ID:       uploadFile.ID,
		Hash:     uploadFile.Hash,
		Name:     uploadFile.Name,
		MimeType: uploadFile.MimeType,
		Time:     uploadFile.Time,
	}

	err = s.Db.WithContext(ctx).Save(uploadModel).Error
	if err != nil {
		return err
	}

	return nil
}
