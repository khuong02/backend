package main

import (
	"context"
	"fmt"
	"github.com/khuong02/backend/cmd/server/config"
	"github.com/khuong02/backend/internal/user/app"
	serviceHttp "github.com/khuong02/backend/internal/user/delivery/http"
	"github.com/khuong02/backend/internal/user/docs"
	"github.com/khuong02/backend/internal/user/migrations"
	"github.com/khuong02/backend/internal/user/repositories"
	_auth "github.com/khuong02/backend/internal/user/usecases/auth"
	_media "github.com/khuong02/backend/internal/user/usecases/media"
	"github.com/khuong02/backend/pkg/database"
	"github.com/khuong02/backend/pkg/flags"
	"github.com/khuong02/backend/pkg/logger"
	"github.com/khuong02/backend/pkg/minio"
	"github.com/soheilhy/cmux"
	"log"
	"net"
	"net/http"
	"time"
)

//	@title		User API
//	@version	1.0

//	@BasePath	/v1/api
//	@schemes	http https

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

// @description	Transaction API.
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// logger
	l, err := logger.NewLogger(cfg.Log.LoggerPath, -4)
	if err != nil {
		log.Println("Get logger failed")
	}

	// database
	psql := database.New(*cfg.Postgres, l)
	psql.GetDNSSchemaDB().Connect()

	defer psql.Disconnect()

	// minio
	minioClient := minio.NewMinioClient(*cfg.Minio, l).Connect()

	// repo
	userRepo := repositories.NewAuth(psql.GetClient)
	mediaRepo := repositories.NewMedia(psql.GetClient)

	// usecases
	auth := _auth.NewAuth(userRepo, l, *cfg)
	media := _media.NewMedia(mediaRepo, l, *cfg, *minioClient)

	service := app.New(l, *cfg, auth, media)

	switch flags.Task {
	case "server":
		executeServer(service, l)
	case "migration-down":
		migration := migrations.New(psql.GetClient(context.Background()), *cfg.Postgres, l)
		migration.Down()
	case "migration-up":
		migration := migrations.New(psql.GetClient(context.Background()), *cfg.Postgres, l)
		migration.Up()
	default:
		executeServer(service, l)
	}
}

func executeServer(service *app.Service, logger *logger.Logger) {
	// cron job
	// job.Init(service)

	// swagger
	docs.SwaggerInfo.Host = service.Cfg.Swagger.URL

	l, err := net.Listen("tcp", fmt.Sprintf("%v:%v", service.Cfg.Http.Host, service.Cfg.Http.Port))
	if err != nil {
		logger.Error("Listen service", "err: ", err)

		return
	}

	m := cmux.New(l)
	httpL := m.Match(cmux.HTTP1Fast())
	errs := make(chan error)

	// http
	go func() {
		httpS := &http.Server{
			Addr:              fmt.Sprintf(":%v", service.Cfg.Http.Port),
			Handler:           serviceHttp.New(service, logger).NewHTTPHandler(),
			ReadHeaderTimeout: 3 * time.Second,
		}
		logger.Info("Server running on: " + fmt.Sprintf("%v:%v", service.Cfg.Http.Host, service.Cfg.Http.Port))

		errs <- httpS.Serve(httpL)
	}()

	go func() {
		errs <- m.Serve()
	}()

	err = <-errs
	if err != nil {
		logger.Error("Start service fail", "err: ", err)
	}
}
