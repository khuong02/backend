package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/khuong02/backend/cmd/server/config"
	"github.com/khuong02/backend/internal/user/app"
	_auth "github.com/khuong02/backend/internal/user/usecases/auth"
	_media "github.com/khuong02/backend/internal/user/usecases/media"
	"github.com/khuong02/backend/pkg/helper"
	"github.com/khuong02/backend/pkg/logger"
	"github.com/khuong02/backend/pkg/my_jwt"
	"net/http"
)

type Route struct {
	Auth  _auth.IAuth
	Media _media.IMedia

	Cfg    config.Config
	logger *logger.Logger
}

func Init(group *gin.RouterGroup, service *app.Service) {
	r := &Route{
		Auth:  service.Auth,
		Media: service.Media,

		Cfg:    service.Cfg,
		logger: service.Logger,
	}

	// user
	group.GET("user/healthcheck", func(c *gin.Context) {
		helper.Success(c, http.StatusOK, "successfully", nil)
	})

	// auth
	group.POST("/user/register", r.Register)
	group.POST("/user/login", r.Login)

	// media
	group.Use(my_jwt.VerifyJWTMiddleware(r.Cfg), my_jwt.AuthenticationJWTType)
	group.POST("/upload", r.Upload)
}
