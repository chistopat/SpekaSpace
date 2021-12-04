package db

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog"
)

const (
	Host       = "127.0.0.1"
	Port       = "5432"
	DriverName = "postgres"
	SSLMode    = "disable"
)

type Config struct {
	Host         string
	Port         string
	DatabaseName string
	User         string
	Password     string
	SslMode      string
	SourceURL    string
	DriverName   string
	Debug        bool
}

func (db *Config) String() string {
	return fmt.Sprintf("Connecting to DB on %s:%s/%s as '%s' ...", db.Host, db.Port, db.DatabaseName, db.User)
}

func (db *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s sslmode=%s user=%s password=%s",
		db.Host, db.Port, db.DatabaseName, db.SslMode, db.User, db.Password,
	)
}

func (db *Config) DefaultDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s sslmode=%s user=%s password=%s",
		db.Host, db.Port, "postgres", db.SslMode, db.User, db.Password,
	)
}

func (db *Config) Copy() *Config {
	return &Config{
		Host:         db.Host,
		Port:         db.Port,
		DatabaseName: db.DatabaseName,
		User:         db.User,
		Password:     db.Password,
		SslMode:      db.SslMode,
		SourceURL:    db.SourceURL,
		DriverName:   db.DriverName,
		Debug:        db.Debug,
	}
}

func GetConnect(cfg *Config, log *zerolog.Logger) (*sql.DB, error) {
	log.Info().Msg(fmt.Sprintf("Connecting to DB on %s:%s/%s as '%s' ...", cfg.Host, cfg.Port, cfg.DatabaseName, cfg.User))
	db, err := sql.Open(cfg.DriverName, cfg.DSN())
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return db, nil
}
