//go:generate mockgen -source=example.go -destination=../mocks/mock_repository.go -package=mocks ExampleQuery

package repositories

import (
	"context"
	"gorm.io/gorm"
)

type userRepo struct {
	getClient func(ctx context.Context) *gorm.DB
}

func NewUser(getClient func(ctx context.Context) *gorm.DB) IUser {
	return &userRepo{getClient: getClient}
}
