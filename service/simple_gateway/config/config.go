package config

import (
	"git.redmadrobot.com/internship/backend/lim-ext/pkg/http"
	"git.redmadrobot.com/internship/backend/lim-ext/pkg/logger"
	"git.redmadrobot.com/internship/backend/lim-ext/pkg/utils"
	"github.com/spf13/viper"
)

const (
	envPrefix                    = "LIM"
	envSimpleGwServerEnvironment = "SIMPLE_GW_SERVER_ENVIRONMENT"
	envSimpleGwServerScheme      = "SIMPLE_GW_SERVER_SCHEME"
	envSimpleGwServerHost        = "SIMPLE_GW_SERVER_HOST"
	envSimpleGwServerPort        = "SIMPLE_GW_SERVER_PORT"
)

func NewViper() *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(envPrefix)
	v.SetDefault(logger.EnvLogLevel, "info")
	v.AutomaticEnv()

	return v
}

type SimpleGWConfig struct {
	Logger            *logger.Config
	Server            *http.Config
	ServerEnvironment utils.ServerEnvironment
}

func NewSimpleGwServiceConfig() *SimpleGWConfig {
	v := NewViper()
	v.SetDefault(envSimpleGwServerEnvironment, utils.DEV)
	v.SetDefault(envSimpleGwServerScheme, http.Scheme)
	v.SetDefault(envSimpleGwServerHost, http.Host)
	v.SetDefault(envSimpleGwServerPort, http.Port)
	return &SimpleGWConfig{
		Logger:            logger.NewLoggerConfig(v),
		Server:            NewServerSimpleGwServiceConfig(v),
		ServerEnvironment: utils.ServerEnvironment(v.GetString(envSimpleGwServerEnvironment)),
	}
}

func NewServerSimpleGwServiceConfig(v *viper.Viper) *http.Config {
	return &http.Config{
		Scheme: v.GetString(envSimpleGwServerScheme),
		Host:   v.GetString(envSimpleGwServerHost),
		Port:   v.GetString(envSimpleGwServerPort),
	}
}
