package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/khuong02/backend/cmd/server/config"
	"github.com/khuong02/backend/internal/user/codeerror"
	"github.com/khuong02/backend/internal/user/dtos"
	"github.com/khuong02/backend/internal/user/payload"
	"github.com/khuong02/backend/internal/user/repositories"
	"github.com/khuong02/backend/pkg/logger"
	"github.com/khuong02/backend/pkg/my_jwt"
	"time"
)

type Auth struct {
	repo   repositories.IAuth
	logger *logger.Logger
	cfg    config.Config
}

func NewAuth(AuthRepo repositories.IAuth, logger *logger.Logger, cfg config.Config) IAuth {
	return &Auth{
		repo:   AuthRepo,
		logger: logger,
		cfg:    cfg,
	}
}

func (uc *Auth) Register(ctx context.Context, req payload.Register) (*dtos.AuthResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	user, err := uc.repo.CreateUser(ctx, req.ToDTO())
	if err != nil {
		return nil, codeerror.ErrRegisterFail(err)
	}

	now := time.Now()
	accessToken, _ := my_jwt.GenerateJWT(my_jwt.MyJWT{
		UserID:   user.ID,
		UserName: user.UserName,
		JWTType:  my_jwt.Authentication,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Duration(uc.cfg.JWT.AccessTokenExpireTime) * time.Second).Unix(),
		},
	}, uc.cfg.JWT.AccessSecretKey)

	return dtos.NewAuthResponse(accessToken), nil
}
