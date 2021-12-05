package config

import (
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/db"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/http"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/logger"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/utils"
	"github.com/spf13/viper"
)

const (
	envServerEnvironment = "SERVER_ENVIRONMENT"
	envServerScheme      = "SERVER_SCHEME"
	envServerHost        = "SERVER_HOST"
	envServerPort        = "SERVER_PORT"

	envDBHost       = "DB_HOST"
	envDBPort       = "DB_PORT"
	envDBName       = "DB_NAME"
	envDBUser       = "DB_USER"
	envDBPassword   = "DB_PASSWORD"
	envDBSslMode    = "DB_SSL_MODE"
	envDBSourceURL  = "DB_SOURCE_URL"
	envDBDriverName = "DB_DRIVER_NAME"
	envDBDebug      = "DB_DEBUG"
)

func NewViper() *viper.Viper {
	v := viper.New()
	v.SetDefault(logger.EnvLogLevel, "info")
	v.AutomaticEnv()

	return v
}

type SpekaSpaceConfig struct {
	Logger            *logger.Config
	Server            *http.Config
	ServerEnvironment utils.ServerEnvironment
	DB                *db.Config
}

func NewSpekaSpaceServiceConfig() *SpekaSpaceConfig {
	v := NewViper()
	//v.SetDefault(envSpekaSpaceServerEnvironment, utils.DEV)
	//v.SetDefault(envSpekaSpaceServerScheme, http.Scheme)
	//v.SetDefault(envSpekaSpaceServerHost, http.Host)
	//v.SetDefault(envSpekaSpaceServerPort, http.Port)
	//
	//v.SetDefault(envSpekaSpaceDBHost, db.Host)
	//v.SetDefault(envSpekaSpaceDBPort, db.Port)
	//v.SetDefault(envSpekaSpaceDBName, "lim")
	//v.SetDefault(envSpekaSpaceDBUser, "postgres")
	//v.SetDefault(envSpekaSpaceDBPassword, "")
	//v.SetDefault(envSpekaSpaceDBSourceURL, "./migrations")
	v.SetDefault(envDBDriverName, db.DriverName)

	return &SpekaSpaceConfig{
		Logger:            logger.NewLoggerConfig(v),
		Server:            NewServerServiceConfig(v),
		ServerEnvironment: utils.ServerEnvironment(v.GetString(envServerEnvironment)),
		DB:                NewDBServiceConfig(v),
	}
}

func NewServerServiceConfig(v *viper.Viper) *http.Config {
	return &http.Config{
		Scheme: v.GetString(envServerScheme),
		Host:   v.GetString(envServerHost),
		Port:   v.GetString(envServerPort),
	}
}

func NewDBServiceConfig(v *viper.Viper) *db.Config {
	return &db.Config{
		Host:         v.GetString(envDBHost),
		Port:         v.GetString(envDBPort),
		DatabaseName: v.GetString(envDBName),
		User:         v.GetString(envDBUser),
		Password:     v.GetString(envDBPassword),
		SslMode:      v.GetString(envDBSslMode),
		SourceURL:    v.GetString(envDBSourceURL),
		DriverName:   v.GetString(envDBDriverName),
		Debug:        v.GetBool(envDBDebug),
	}
}
