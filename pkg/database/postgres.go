package database

import (
	"context"
	"fmt"
	"github.com/khuong02/backend/pkg/logger"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type PostgresConfig struct {
	Host   string `env-required:"true" yaml:"HOST"    env:"POSTGRES_DB_HOST"`
	Port   int    `env-required:"true" yaml:"PORT"    env:"POSTGRES_DB_PORT"`
	DB     string `env-required:"true" yaml:"NAME"    env:"POSTGRES_DB_NAME"`
	User   string `env-required:"true" yaml:"USER"    env:"POSTGRES_DB_USER"`
	Pass   string `env-required:"true" yaml:"PASS"    env:"POSTGRES_DB_PASS"`
	Schema string `env-required:"true" yaml:"SCHEMA"    env:"POSTGRES_DB_SCHEMA"`

	DBMaxIdleConns  int `yaml:"MAX_IDLE_CONNS" env:"POSTGRES_DB_MAX_IDLE_CONNS"`
	DBMaxOpenConns  int `yaml:"MAX_OPEN_CONNS" env:"POSTGRES_DB_MAX_OPEN_CONNS"`
	CountRetryTx    int `yaml:"TX_RETRY_COUNT" env:"POSTGRES_DB_TX_RETRY_COUNT"`
	ConnMaxLifeTime int `yaml:"CONN_MAX_LIFE_TIME" env:"POSTGRES_DB_CONN_MAX_LIFE_TIME"`

	DirMigration string `yaml:"DIR_MIGRATION" env:"POSTGRES_DB_DIR_MIGRATION"`
}

type Postgres struct {
	db        *gorm.DB
	dns       string
	schemaStr string
	cfg       PostgresConfig
	logger    *logger.Logger
}

func New(cfg PostgresConfig, logger *logger.Logger) *Postgres {
	return &Postgres{
		cfg:    cfg,
		logger: logger,
	}
}

func (pg *Postgres) Connect() {
	var err error

	pg.db, err = gorm.Open(postgres.Open(pg.dns), &gorm.Config{
		//Logger: cengormlogger.NewLogger(pg.logger,
		//	int(slog.LevelDebug)),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   pg.schemaStr,
			SingularTable: false,
		},
	})
	if err != nil {
		pg.logger.Error("Connect postgres fail", "err: ", err)

		return
	}

	if err := pg.db.Use(otelgorm.NewPlugin()); err != nil {
		pg.logger.Error("otel gorm fail", "err:", err)

		os.Exit(1)
	}
	pg.db = pg.db.Debug()

	rawDB, _ := pg.db.DB()
	rawDB.SetMaxIdleConns(pg.cfg.DBMaxIdleConns)
	rawDB.SetMaxOpenConns(pg.cfg.DBMaxOpenConns)
	rawDB.SetConnMaxLifetime(time.Minute * time.Duration(pg.cfg.ConnMaxLifeTime))

	err = rawDB.Ping()
	if err != nil {
		pg.logger.Error("Ping postgres fail", "err: ", err)

		os.Exit(1)
	}

	pg.logger.Info("Connect postgres successfully!!")
}

func (pg *Postgres) GetClient(ctx context.Context) *gorm.DB {
	return pg.db.WithContext(ctx).Session(&gorm.Session{})
}

func (pg *Postgres) GetDNSSchemaDB() *Postgres {
	pg.dns = fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		pg.cfg.User,
		pg.cfg.Pass,
		pg.cfg.Host,
		pg.cfg.Port,
		pg.cfg.DB,
	)
	pg.schemaStr = pg.cfg.Schema

	return pg
}

func (pg *Postgres) Disconnect() {
	postgresDB, err := pg.db.DB()
	if err != nil {
		pg.logger.Error("Disconnect postgres fail", "err: ", err)

		os.Exit(1)
	}

	if pg.db != nil {
		_ = postgresDB.Close()
	}

	pg.logger.Info("Disconnect postgres successfully!!")
}
