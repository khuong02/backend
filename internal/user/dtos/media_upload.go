package dtos

import (
	"github.com/google/uuid"
)

type MediaUploadDTO struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Path     string    `json:"path"`
	UploadBy uuid.UUID `json:"upload_by"`
}

func NewMediaUploadDTO(userUpload uuid.UUID, name, path string) *MediaUploadDTO {
	return &MediaUploadDTO{
		ID:       uuid.New(),
		Name:     name,
		Path:     path,
		UploadBy: userUpload,
	}
}
