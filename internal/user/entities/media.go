package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Media struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	UploadBy  uuid.UUID      `json:"upload_by"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (Media) GetTableName() string {
	return "media"
}
