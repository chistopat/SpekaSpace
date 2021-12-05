package db

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	entSql "entgo.io/ent/dialect/sql"
	mgr "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// Need to work with migration files.
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog"
)

func GetDriver(cfg *Config, log *zerolog.Logger) (*entSql.Driver, error) {
	db, err := GetConnect(cfg, log)
	if err != nil {
		return nil, err
	}
	return entSql.OpenDB(cfg.DriverName, db), nil
}

func Migrate(cfg *Config, log *zerolog.Logger) error {
	db, err := sql.Open(cfg.DriverName, cfg.DSN())
	if err != nil {
		log.Error().Err(err).Msg(err.Error())

		return err
	}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Info().Str("Dir", dir).Msg(fmt.Sprintf("Start applying DB migrations from source %s", cfg.SourceURL))
	if _, err := os.Stat(cfg.SourceURL); !os.IsNotExist(err) {
		err = applyMigrations(db, cfg, log)
		if err != nil {
			log.Error().Err(err).Msg(err.Error())

			return err
		}
	} else {
		log.Info().Msg(fmt.Sprintf("not found migrations directory: %s", err.Error()))
	}

	return nil
}

func getMigrate(db *sql.DB, cfg *Config, log *zerolog.Logger) (*mgr.Migrate, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("cannot get driverName: %w", err)
	}

	source := "file://" + cfg.SourceURL
	m, err := mgr.NewWithDatabaseInstance(source, cfg.DatabaseName, driver)
	if err != nil {
		return nil, fmt.Errorf("cannot get migrate instance: %w", err)
	}
	m.Log = MigrateLogger{
		log: log,
		cfg: cfg,
	}
	return m, nil
}

func applyMigrations(db *sql.DB, cfg *Config, log *zerolog.Logger) error {
	isEmptySource, err := isEmpty(cfg.SourceURL)
	if err != nil {
		return fmt.Errorf("not found source: %w", err)
	}
	if isEmptySource {
		return nil
	}
	m, err := getMigrate(db, cfg, log)
	if err != nil {
		return fmt.Errorf("cannot get migration: %w", err)
	}
	err = m.Up()
	if err != nil {
		if !errors.Is(err, mgr.ErrNoChange) {
			return fmt.Errorf("migration error: %w", err)
		}
	}

	return nil
}

// isEmpty checks if there are files in the directory
func isEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}

	return false, err // Either not empty or error, suits both cases
}

type MigrateLogger struct {
	log *zerolog.Logger
	cfg *Config
}

func (m MigrateLogger) Printf(format string, v ...interface{}) {
	msg := fmt.Sprintf("Migration: %s", strings.TrimRight(format, "\n"))
	m.log.Info().Msg(fmt.Sprintf(msg, v...))
}

func (m MigrateLogger) Verbose() bool {
	return m.cfg.Debug
}
