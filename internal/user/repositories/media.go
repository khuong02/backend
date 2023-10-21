package repositories

import (
	"context"
	"github.com/khuong02/backend/internal/user/dtos"
	entities "github.com/khuong02/backend/internal/user/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mediaRepo struct {
	getClient func(ctx context.Context) *gorm.DB
}

func NewMedia(getClient func(ctx context.Context) *gorm.DB) IMedia {
	return &mediaRepo{getClient: getClient}
}

func (pg *mediaRepo) Insert(ctx context.Context, mediaUpload dtos.MediaUploadDTO) error {
	return pg.getClient(ctx).Table(entities.Media{}.GetTableName()).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}, {Name: "path"}},
			DoUpdates: clause.AssignmentColumns([]string{"upload_by"}),
		}).Create(&mediaUpload).Error
}
