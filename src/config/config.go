package config

import (
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/db"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/http"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/logger"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/utils"
	"github.com/spf13/viper"
)

const (
	envPrefix                      = "LIM"
	envSpekaSpaceServerEnvironment = "SPEKASPACE_SERVER_ENVIRONMENT"
	envSpekaSpaceServerScheme      = "SPEKASPACE_SERVER_SCHEME"
	envSpekaSpaceServerHost        = "SPEKASPACE_SERVER_HOST"
	envSpekaSpaceServerPort        = "SPEKASPACE_SERVER_PORT"

	envSpekaSpaceDBHost       = "SPEKASPACE_DB_HOST"
	envSpekaSpaceDBPort       = "SPEKASPACE_DB_PORT"
	envSpekaSpaceDBName       = "SPEKASPACE_DB_NAME"
	envSpekaSpaceDBUser       = "SPEKASPACE_DB_USER"
	envSpekaSpaceDBPassword   = "SPEKASPACE_DB_PASSWORD"
	envSpekaSpaceDBSslMode    = "SPEKASPACE_DB_SSL_MODE"
	envSpekaSpaceDBSourceURL  = "SPEKASPACE_DB_SOURCE_URL"
	envSpekaSpaceDBDriverName = "SPEKASPACE_DB_DRIVER_NAME"
	envSpekaSpaceDBDebug      = "SPEKASPACE_DB_DEBUG"
)

func NewViper() *viper.Viper {
	v := viper.New()
	//v.SetEnvPrefix(envPrefix)
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
	v.SetDefault(envSpekaSpaceServerEnvironment, utils.DEV)
	v.SetDefault(envSpekaSpaceServerScheme, http.Scheme)
	v.SetDefault(envSpekaSpaceServerHost, http.Host)
	v.SetDefault(envSpekaSpaceServerPort, http.Port)

	v.SetDefault(envSpekaSpaceDBHost, db.Host)
	v.SetDefault(envSpekaSpaceDBPort, db.Port)
	v.SetDefault(envSpekaSpaceDBName, "lim")
	v.SetDefault(envSpekaSpaceDBUser, "postgres")
	v.SetDefault(envSpekaSpaceDBPassword, "")
	v.SetDefault(envSpekaSpaceDBSslMode, db.SSLMode)
	v.SetDefault(envSpekaSpaceDBSourceURL, "./migrations")
	v.SetDefault(envSpekaSpaceDBDriverName, db.DriverName)
	v.SetDefault(envSpekaSpaceDBDebug, false)

	return &SpekaSpaceConfig{
		Logger:            logger.NewLoggerConfig(v),
		Server:            NewServerSpekaSpaceServiceConfig(v),
		ServerEnvironment: utils.ServerEnvironment(v.GetString(envSpekaSpaceServerEnvironment)),
		DB:                NewDBSpekaSpaceServiceConfig(v),
	}
}

func NewServerSpekaSpaceServiceConfig(v *viper.Viper) *http.Config {
	return &http.Config{
		Scheme: v.GetString(envSpekaSpaceServerScheme),
		Host:   v.GetString(envSpekaSpaceServerHost),
		Port:   v.GetString(envSpekaSpaceServerPort),
	}
}

func NewDBSpekaSpaceServiceConfig(v *viper.Viper) *db.Config {
	return &db.Config{
		Host:         v.GetString(envSpekaSpaceDBHost),
		Port:         v.GetString(envSpekaSpaceDBPort),
		DatabaseName: v.GetString(envSpekaSpaceDBName),
		User:         v.GetString(envSpekaSpaceDBUser),
		Password:     v.GetString(envSpekaSpaceDBPassword),
		SslMode:      v.GetString(envSpekaSpaceDBSslMode),
		SourceURL:    v.GetString(envSpekaSpaceDBSourceURL),
		DriverName:   v.GetString(envSpekaSpaceDBDriverName),
		Debug:        v.GetBool(envSpekaSpaceDBDebug),
	}
}
