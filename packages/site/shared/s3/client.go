package s3

import (
	"context"
	"time"

	"github.com/4strodev/4stroblog/site/shared/config"
	"github.com/minio/minio-go/v7"
)

const (
	UPLOADS_BUCKET string = "uploads"
)

func NewS3Client(config config.Config) (*minio.Client, error) {
	client, err := minio.New(config.Storage.S3.Url, nil)
	if err != nil {
		return nil, err
	}

	var ctx context.Context
	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	err = client.MakeBucket(ctx, "uploads", minio.MakeBucketOptions{})
	if err != nil {
		return nil, err
	}

	return nil, err
}
