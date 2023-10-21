package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/khuong02/backend/cmd/server/config"
	"github.com/khuong02/backend/internal/user/app"
	_user "github.com/khuong02/backend/internal/user/usecases/user"
	"github.com/khuong02/backend/pkg/helper"
	"github.com/khuong02/backend/pkg/logger"
	"net/http"
)

type Route struct {
	User _user.IUser

	Cfg    config.Config
	logger *logger.Logger
}

func Init(group *gin.RouterGroup, service *app.Service) {
	_ = &Route{
		User: service.User,

		Cfg:    service.Cfg,
		logger: service.Logger,
	}

	// user
	group.GET("user/healthcheck", func(c *gin.Context) {
		helper.Success(c, http.StatusOK, "successfully", nil)
	})
}
