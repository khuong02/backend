package user

import (
	"github.com/khuong02/backend/internal/user/repositories"
	"github.com/khuong02/backend/pkg/logger"
)

type User struct {
	repo   repositories.IUser
	logger *logger.Logger
}

func NewUser(userRepo repositories.IUser, logger *logger.Logger) IUser {
	return &User{
		repo:   userRepo,
		logger: logger,
	}
}
