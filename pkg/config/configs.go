package config

import (
	"github.com/khuong02/backend/pkg/stage"
)

type (
	App struct {
		Name    string          `env-required:"true" yaml:"NAME"    env:"APP_NAME"`
		Version string          `env-required:"true" yaml:"VERSION" env:"APP_VERSION"`
		Stage   stage.StageType `env-required:"true" yaml:"STAGE"    env:"APP_STAGE"`
	}
	Http struct {
		Port string `env-required:"true" yaml:"PORT"    env:"HTTP_PORT"`
		Host string `env-required:"true" yaml:"HOST"    env:"HTTP_HOST"`
	}
	Swagger struct {
		URL string `env-required:"true" yaml:"URL"    env:"SWAGGER_URL"`
	}
	Log struct {
		LoggerPath string `env-required:"true" yaml:"PATH_LOGGER"    env:"PATH_LOGGER"`
	}
)
