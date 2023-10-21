package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	configs "github.com/khuong02/backend/pkg/config"
	"github.com/khuong02/backend/pkg/database"
	"github.com/khuong02/backend/pkg/flags"
	"log"
	"os"
)

type Config struct {
	configs.App     `yaml:"APP"`
	configs.Http    `yaml:"HTTP"`
	configs.Log     `yaml:"LOG"`
	configs.Swagger `yaml:"SWAGGER"`
	JWT             JWT                      `yaml:"JWT"`
	Postgres        *database.PostgresConfig `yaml:"POSTGRES_DB"`
}

type Migration struct {
	Dir string `mapstructure:"DIR_MIGRATION"`
}

type JWT struct {
	AccessSecretKey       string `yaml:"ACCESS_SECRET_KEY"    env:"JWT_ACCESS_SECRET_KEY"`
	AccessTokenExpireTime int    `yaml:"ACCESS_TOKEN_EXPIRE_TIME" env:"JWT_ACCESS_TOKEN_EXPIRE_TIME"` // unit second
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load()
	flags.GetFlag()
	cfg := &Config{}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// debug
	fmt.Println(dir + "/config.yaml")

	err = cleanenv.ReadConfig(dir+"/config.yaml", cfg)
	if err != nil {
		err = cleanenv.ReadEnv(cfg)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
