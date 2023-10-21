package migrations

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	db "github.com/khuong02/backend/pkg/database"
	"github.com/khuong02/backend/pkg/logger"

	// Register using Golang migrate.
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Migration struct {
	DBName    string
	SourceURL string
	DB        *gorm.DB
	Logger    *logger.Logger
}

func New(db *gorm.DB, config db.PostgresConfig, logger *logger.Logger) *Migration {
	return &Migration{
		DBName:    config.DB,
		SourceURL: config.DirMigration,
		DB:        db,
		Logger:    logger,
	}
}

func (mig *Migration) initDBInstance() (*migrate.Migrate, error) {
	getDB, err := mig.DB.DB()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		mig.Logger.Error(err.Error())

		return nil, err
	}

	driver, err := postgres.WithInstance(getDB, &postgres.Config{MigrationsTable: "migration"})
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		mig.Logger.Error(err.Error())

		return nil, err
	}

	return migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%v", mig.SourceURL), mig.DBName, driver)
}

func (mig *Migration) Up() {
	m, err := mig.initDBInstance()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		mig.Logger.Error(err.Error())

		return
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		mig.Logger.Error(err.Error())

		return
	}

	mig.Logger.Info("Up done!")
}

func (mig *Migration) Down() {
	m, err := mig.initDBInstance()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		mig.Logger.Error(err.Error())

		return
	}

	err = m.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		mig.Logger.Error(err.Error())

		return
	}

	mig.Logger.Info("Down done!")
}
