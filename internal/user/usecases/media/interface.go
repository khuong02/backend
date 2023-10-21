package media

import (
	"context"
	"github.com/khuong02/backend/internal/user/dtos"
	"github.com/khuong02/backend/internal/user/payload"
)

type IMedia interface {
	UploadMedia(ctx context.Context, req payload.MediaUpload) (*dtos.MediaUploadResp, error)
}
