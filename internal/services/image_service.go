package services

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime"
	"path/filepath"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type ImageService struct {
	client     *minio.Client
	bucketName string
}

func NewImageService(cfg config.MinIOConfig) (*ImageService, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Login, cfg.Password, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %w", err)
	}

	// Создаём бакет, если не существует
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, cfg.BucketName)
	if err != nil {
		return nil, fmt.Errorf("bucket check failed: %w", err)
	}
	if !exists {
		err = client.MakeBucket(ctx, cfg.BucketName, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("bucket creation failed: %w", err)
		}
	}

	return &ImageService{
		client:     client,
		bucketName: cfg.BucketName,
	}, nil
}

func (s *ImageService) UploadObject(ctx context.Context, key string, data []byte) error {
	_, err := s.client.PutObject(
		ctx,
		s.bucketName,
		key,
		bytes.NewReader(data),
		int64(len(data)),
		minio.PutObjectOptions{},
	)
	return err
}

func (s *ImageService) GetPresignedURL(ctx context.Context, key string) (string, error) {
	url, err := s.client.PresignedGetObject(ctx, s.bucketName, key, 15*time.Minute, nil)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}
	return url.String(), nil
}

func (s *ImageService) DeleteObject(ctx context.Context, key string) error {
	return s.client.RemoveObject(ctx, s.bucketName, key, minio.RemoveObjectOptions{})
}

func (s *ImageService) GetFileByMinioKey(ctx context.Context, minioKey, originalFilename string) ([]byte, string, error) {
	contentType := mime.TypeByExtension(filepath.Ext(originalFilename))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	reader, err := s.client.GetObject(ctx, s.bucketName, minioKey, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", errors.New("failed to retrieve file from storage")
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, "", errors.New("failed to read file")
	}

	return data, contentType, nil
}
