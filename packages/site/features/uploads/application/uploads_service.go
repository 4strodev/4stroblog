package application

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"

	"github.com/4strodev/4stroblog/site/features/uploads/domain"
	"github.com/4strodev/4stroblog/site/shared/db/models"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

func NewUploadsService(db *gorm.DB, s3 *minio.Client) *UploadsService {
	return &UploadsService{
		Db: db,
		S3: s3,
	}
}

type UploadsService struct {
	Db *gorm.DB
	S3 *minio.Client
}

func (s *UploadsService) SaveFile() error {
	var upload domain.Upload
	if upload.Content == nil {
		return errors.New("cannot upload an upload with no content")
	}

	var buffer bytes.Buffer
	var contentReader = io.TeeReader(upload.Content, &buffer)

	content, err := io.ReadAll(contentReader)
	if err != nil {
		return err
	}
	checksum := sha256.Sum256(content)
	upload.Hash = hex.Dump(checksum[:])

	var ctx = context.Background()
	_, err = s.S3.PutObject(ctx, "uploads", upload.Hash, &buffer, int64(buffer.Len()), minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	uploadModel := models.Upload{
		ID:       upload.ID,
		Hash:     upload.Hash,
		Name:     upload.Name,
		MimeType: upload.MimeType,
	}

	s.Db.Save()
}
