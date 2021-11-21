package main

import (
	"fmt"
	"net/http"
	"os"

	"git.redmadrobot.com/internship/backend/lim-ext/pkg/logger"
	"git.redmadrobot.com/internship/backend/lim-ext/service/simple_gateway"

	"git.redmadrobot.com/internship/backend/lim-ext/service/simple_gateway/config"
	"github.com/rs/zerolog/log"

	"git.redmadrobot.com/internship/backend/lim-ext/service/simple_gateway/generated"
)

func main() {
	if err := run(); err != nil {
		log.Error().Err(err).Msg("shutting down")
		os.Exit(1)
	}
}

func run() error {
	// Configuration
	serviceConfig := config.NewSimpleGwServiceConfig()
	address := fmt.Sprintf("%s:%s", serviceConfig.Server.Host, serviceConfig.Server.Port)

	// Logging
	serviceLogger, err := logger.NewLogger(serviceConfig.Logger)
	if err != nil {
		return err
	}

	server := simple_gateway.NewServer(serviceLogger, serviceConfig)
	options := simple_gateway.NewServerOptions(&server)
	router := generated.HandlerWithOptions(server, options)

	serviceLogger.Info().Msgf("Start Simple GateWay server: %s", address)

	err = http.ListenAndServe(address, router)
	if err != nil {
		serviceLogger.Error().Msg("Error serving http")
	}

	return err
}
