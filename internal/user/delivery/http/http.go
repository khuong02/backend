package http

import (
	"fmt"
	"github.com/khuong02/backend/internal/user/app"
	user "github.com/khuong02/backend/internal/user/routers"
	"github.com/khuong02/backend/pkg/logger"
	"net/http"
	"time"

	swaggerfiles "github.com/swaggo/files"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServiceHTTP struct {
	service *app.Service
	// grpc    *app.GRPC
	logger *logger.Logger
}

func New(service *app.Service, logger *logger.Logger) *ServiceHTTP {
	return &ServiceHTTP{
		service: service,
		// grpc:    grpc,
		logger: logger,
	}
}

func (h *ServiceHTTP) NewHTTPHandler() *gin.Engine {
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%q\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodPut,
			http.MethodPatch,
			http.MethodOptions,
		},
		AllowHeaders:           []string{"Origin", "Authorization", "Content-Type", "token"},
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             true,
	}))
	r.Use(gin.Recovery())

	// API docs
	if !h.service.Cfg.Stage.IsProd() {
		r.GET("docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	user.Init(r.Group("/v1/api"), h.service)

	return r
}
