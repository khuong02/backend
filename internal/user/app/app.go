package app

import (
	"github.com/khuong02/backend/cmd/server/config"
	_auth "github.com/khuong02/backend/internal/user/usecases/auth"
	"github.com/khuong02/backend/pkg/logger"
)

type Service struct {
	// usecases
	Auth _auth.IAuth

	// plugins
	Logger *logger.Logger
	Cfg    config.Config
}

func New(
	// plugins
	logger *logger.Logger,
	cfg config.Config,

	// usecases
	auth _auth.IAuth,
) *Service {
	return &Service{
		Auth: auth,

		Logger: logger,
		Cfg:    cfg,
	}
}
