package media

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/khuong02/backend/cmd/server/config"
	"github.com/khuong02/backend/internal/user/codeerror"
	"github.com/khuong02/backend/internal/user/constants"
	"github.com/khuong02/backend/internal/user/dtos"
	"github.com/khuong02/backend/internal/user/payload"
	"github.com/khuong02/backend/internal/user/repositories"
	"github.com/khuong02/backend/pkg/logger"
	"github.com/khuong02/backend/pkg/minio"
)

type Media struct {
	repo repositories.IMedia
	// plugins
	logger *logger.Logger
	cfg    config.Config

	// clients
	minioClient minio.MinioClient
}

func NewMedia(
	repo repositories.IMedia,
	logger *logger.Logger,
	cfg config.Config,
	minioClient minio.MinioClient,
) IMedia {
	return &Media{
		repo:        repo,
		logger:      logger,
		cfg:         cfg,
		minioClient: minioClient,
	}
}

func (uc *Media) UploadMedia(ctx context.Context, req payload.MediaUpload) (*dtos.MediaUploadResp, error) {
	userID := ctx.Value("user_id")

	info, err := uc.minioClient.UploadMedia(ctx, constants.Bucket, req.Media)
	if err != nil {
		uc.logger.Error("upload fail", "err:", err)

		return nil, err
	}

	path := fmt.Sprintf("%v/%v/%v", uc.cfg.Minio.PublicEndpoint, info.Bucket, req.Media.Filename)

	if err := uc.repo.Insert(ctx, *dtos.NewMediaUploadDTO(userID.(uuid.UUID), req.Media.Filename, path)); err != nil {
		uc.logger.Error("insert fail", "err:", err)

		return nil, codeerror.ErrInsertMediaToDB(err)
	}

	return dtos.NewMediaUploadResp(path), nil
}
