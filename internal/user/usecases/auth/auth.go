package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/khuong02/backend/cmd/server/config"
	"github.com/khuong02/backend/internal/user/codeerror"
	"github.com/khuong02/backend/internal/user/dtos"
	"github.com/khuong02/backend/internal/user/entities"
	"github.com/khuong02/backend/internal/user/payload"
	"github.com/khuong02/backend/internal/user/repositories"
	"github.com/khuong02/backend/pkg/logger"
	"github.com/khuong02/backend/pkg/my_jwt"
	"github.com/khuong02/backend/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
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

func (uc *Auth) accessToken(user *entities.User) string {
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

	return accessToken
}

func (uc *Auth) Register(ctx context.Context, req payload.Register) (*dtos.AuthResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	user, err := uc.repo.CreateUser(ctx, req.ToDTO())
	if err != nil {
		return nil, codeerror.ErrRegisterFail(err)
	}

	return dtos.NewAuthResponse(uc.accessToken(user)), nil
}

func (uc *Auth) Login(ctx context.Context, req payload.Login) (*dtos.AuthResponse, error) {
	if err := req.Validate(); err != nil {

		return nil, err
	}

	user, err := uc.repo.FindByUserName(ctx, req.UserName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, codeerror.ErrNotExistUser(err)
		}

		return nil, codeerror.ErrLoginFailed(err)
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, codeerror.ErrWrongPassword(errors.New("Wrong Password"))
	}

	return dtos.NewAuthResponse(uc.accessToken(user)), nil
}
