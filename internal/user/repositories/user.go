//go:generate mockgen -source=example.go -destination=../mocks/mock_repository.go -package=mocks ExampleQuery

package repositories

import (
	"context"
	"github.com/khuong02/backend/internal/user/dtos"
	"github.com/khuong02/backend/internal/user/entities"
	"gorm.io/gorm"
)

type authRepo struct {
	getClient func(ctx context.Context) *gorm.DB
}

func NewAuth(getClient func(ctx context.Context) *gorm.DB) IAuth {
	return &authRepo{getClient: getClient}
}

func (pg *authRepo) CreateUser(ctx context.Context, userDTO dtos.CreateUser) (*entities.User, error) {
	var user entities.User
	err := pg.getClient(ctx).Table(entities.User{}.GetTableName()).
		Create(&userDTO).Where("user_name = ?", userDTO.UserName).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (pg *authRepo) FindByUserName(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User
	err := pg.getClient(ctx).Table(entities.User{}.GetTableName()).
		Where("user_name = ?", username).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
