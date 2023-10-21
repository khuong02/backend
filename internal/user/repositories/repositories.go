package repositories

import (
	"context"
	"github.com/khuong02/backend/internal/user/dtos"
	"github.com/khuong02/backend/internal/user/entities"
)

type (
	IAuth interface {
		CreateUser(ctx context.Context, user dtos.CreateUser) (*entities.User, error)
		FindByUserName(ctx context.Context, username string) (*entities.User, error)
	}
)
