package auth

import (
	"context"
	"github.com/khuong02/backend/internal/user/dtos"
	"github.com/khuong02/backend/internal/user/payload"
)

type IAuth interface {
	Register(ctx context.Context, req payload.Register) (*dtos.AuthResponse, error)
}
