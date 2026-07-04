package application

import (
	"context"
	"net/url"
	"time"

	"github.com/4strodev/4stroblog/site/features/uploads/domain"
	"github.com/4strodev/4stroblog/site/shared/domain/domainerrors"
	"github.com/4strodev/4stroblog/site/shared/s3"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

func NewUploadsService(db *gorm.DB, s3 *minio.Client, uploadsRepository domain.UploadsRepository) *UploadsService {
	return &UploadsService{
		Db:                db,
		ObjectStorage:     s3,
		UploadsRepository: uploadsRepository,
	}
}

type UploadsService struct {
	Db                *gorm.DB
	ObjectStorage     *minio.Client
	UploadsRepository domain.UploadsRepository
}

func createPolicy(upload domain.Upload) (*minio.PostPolicy, error) {
	policy := minio.NewPostPolicy()
	err := policy.SetBucket(s3.UPLOADS_BUCKET)
	if err != nil {
		return nil, domainerrors.Errorf(domainerrors.RUNTIME,
			"cannot set bucket to post policy: %w",
			err)
	}
	err = policy.SetExpires(time.Now().UTC().Add(time.Second * 30))
	if err != nil {
		return nil, domainerrors.Errorf(domainerrors.RUNTIME,
			"cannot set expiration to post policy: %w",
			err)
	}
	err = policy.SetContentType(upload.MimeType)
	if err != nil {
		return nil, domainerrors.Errorf(domainerrors.RUNTIME,
			"cannot set content type to post policy: %w",
			err)
	}
	err = policy.SetKey(upload.ID.String())
	if err != nil {
		return nil, domainerrors.Errorf(domainerrors.RUNTIME,
			"cannot set key to post policy: %w",
			err)
	}

	return policy, nil
}

// CreateUpload saves the passed upload in case it doesn't exists.
// It returns the post policy with the required form data fields for the upload.
// If upload exists returns [domainerrors.DATA_CONFLICT]
func (s *UploadsService) CreateUpload(ctx context.Context, uploadFile domain.Upload) (u *url.URL, formData map[string]string, err error) {
	if uploadFile.Time.IsZero() {
		uploadFile.Time = time.Now()
	}

	// Create pre-signed policy only if upload doesn't exists
	_, err = s.UploadsRepository.FindByID(ctx, uploadFile.ID)
	_, isNotFound := domainerrors.Is(err, domainerrors.ENTITY_NOT_FOUND)
	if !isNotFound && err != nil {
		return
	}

	var policy *minio.PostPolicy
	policy, err = createPolicy(uploadFile)
	if err != nil {
		return
	}

	u, formData, err = s.ObjectStorage.PresignedPostPolicy(ctx, policy)
	return
}
