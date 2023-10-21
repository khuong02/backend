package app

import (
	"github.com/khuong02/backend/cmd/server/config"
	_auth "github.com/khuong02/backend/internal/user/usecases/auth"
	_media "github.com/khuong02/backend/internal/user/usecases/media"
	"github.com/khuong02/backend/pkg/logger"
)

type Service struct {
	// usecases
	Auth  _auth.IAuth
	Media _media.IMedia

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
	media _media.IMedia,
) *Service {
	return &Service{
		Auth:  auth,
		Media: media,

		Logger: logger,
		Cfg:    cfg,
	}
}
