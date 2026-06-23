package minio

import (
	"bytes"
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type StorageClient struct {
	client *minio.Client
	bucket string
}

func NewStorageClient(endpoint, accessKey, secretKey, bucket string, useSSL bool) (*StorageClient, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return &StorageClient{
		client: minioClient,
		bucket: bucket,
	}, nil
}

// UploadFile uploads a byte array to MinIO and returns the object key/path
func (s *StorageClient) UploadFile(ctx context.Context, objectName string, data []byte, contentType string) (string, error) {
	reader := bytes.NewReader(data)
	_, err := s.client.PutObject(ctx, s.bucket, objectName, reader, int64(len(data)), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload to minio: %w", err)
	}

	return objectName, nil
}
