package app

import (
	"github.com/khuong02/backend/cmd/server/config"
	_user "github.com/khuong02/backend/internal/user/usecases/user"
	"github.com/khuong02/backend/pkg/logger"
)

type Service struct {
	// usecases
	User _user.IUser

	// plugins
	Logger *logger.Logger
	Cfg    config.Config
}

func New(
	// plugins
	logger *logger.Logger,
	cfg config.Config,

	// usecases
	user _user.IUser,
) *Service {
	return &Service{
		User: user,

		Logger: logger,
		Cfg:    cfg,
	}
}
