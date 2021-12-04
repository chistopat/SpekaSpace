package config

import (
	"git.redmadrobot.com/internship/backend/lim-ext/pkg/http"
	"git.redmadrobot.com/internship/backend/lim-ext/pkg/logger"
	"git.redmadrobot.com/internship/backend/lim-ext/pkg/utils"
	"github.com/spf13/viper"
)

const (
	envPrefix                      = "LIM"
	envSpekaSpaceServerEnvironment = "SPEKASPACE_SERVER_ENVIRONMENT"
	envSpekaSpaceServerScheme      = "SPEKASPACE_SERVER_SCHEME"
	envSpekaSpaceServerHost        = "SPEKASPACE_SERVER_HOST"
	envSpekaSpaceServerPort        = "SPEKASPACE_SERVER_PORT"
)

func NewViper() *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(envPrefix)
	v.SetDefault(logger.EnvLogLevel, "info")
	v.AutomaticEnv()

	return v
}

type SpekaSpaceConfig struct {
	Logger            *logger.Config
	Server            *http.Config
	ServerEnvironment utils.ServerEnvironment
}

func NewSpekaSpaceServiceConfig() *SpekaSpaceConfig {
	v := NewViper()
	v.SetDefault(envSpekaSpaceServerEnvironment, utils.DEV)
	v.SetDefault(envSpekaSpaceServerScheme, http.Scheme)
	v.SetDefault(envSpekaSpaceServerHost, http.Host)
	v.SetDefault(envSpekaSpaceServerPort, http.Port)
	return &SpekaSpaceConfig{
		Logger:            logger.NewLoggerConfig(v),
		Server:            NewServerSpekaSpaceServiceConfig(v),
		ServerEnvironment: utils.ServerEnvironment(v.GetString(envSpekaSpaceServerEnvironment)),
	}
}

func NewServerSpekaSpaceServiceConfig(v *viper.Viper) *http.Config {
	return &http.Config{
		Scheme: v.GetString(envSpekaSpaceServerScheme),
		Host:   v.GetString(envSpekaSpaceServerHost),
		Port:   v.GetString(envSpekaSpaceServerPort),
	}
}
