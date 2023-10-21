package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/khuong02/backend/cmd/server/config"
	"github.com/khuong02/backend/internal/user/app"
	_auth "github.com/khuong02/backend/internal/user/usecases/auth"
	"github.com/khuong02/backend/pkg/helper"
	"github.com/khuong02/backend/pkg/logger"
	"net/http"
)

type Route struct {
	Auth _auth.IAuth

	Cfg    config.Config
	logger *logger.Logger
}

func Init(group *gin.RouterGroup, service *app.Service) {
	r := &Route{
		Auth: service.Auth,

		Cfg:    service.Cfg,
		logger: service.Logger,
	}

	// user
	group.GET("user/healthcheck", func(c *gin.Context) {
		helper.Success(c, http.StatusOK, "successfully", nil)
	})

	group.POST("/user/register", r.Register)
	group.POST("/user/login", r.Login)
}
