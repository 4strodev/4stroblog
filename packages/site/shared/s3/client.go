package s3

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/4strodev/4stroblog/site/shared/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	UPLOADS_BUCKET string = "uploads"
)

func NewS3Client(config config.Config) (*minio.Client, error) {
	url, err := url.Parse(config.Storage.S3.Url)
	if err != nil {
		return nil, fmt.Errorf("error parsing storage url: %w", err)
	}
	pass, _ := url.User.Password()
	client, err := minio.New(url.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(url.User.Username(), pass, ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating minio client: %s", err)
	}

	var ctx context.Context
	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	exists, err := client.BucketExists(ctx, config.Storage.S3.Bucket)
	if err != nil {
		return nil, err
	}

	if !exists {
		err = client.MakeBucket(ctx, config.Storage.S3.Bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
	}

	return client, err
}
