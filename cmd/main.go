package main

import (
	"fmt"
	"git.redmadrobot.com/internship/backend/lim-ext/src"
	"git.redmadrobot.com/internship/backend/lim-ext/src/config"
	"git.redmadrobot.com/internship/backend/lim-ext/src/generated"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/logger"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

func main() {
	if err := run(); err != nil {
		log.Error().Err(err).Msg("shutting down")
		os.Exit(1)
	}
}

func run() error {
	// Configuration
	serviceConfig := config.NewSpekaSpaceServiceConfig()
	address := fmt.Sprintf("%s:%s", serviceConfig.Server.Host, serviceConfig.Server.Port)

	// Logging
	serviceLogger, err := logger.NewLogger(serviceConfig.Logger)
	if err != nil {
		return err
	}

	server := speka_space.NewServer(serviceLogger, serviceConfig)
	options := speka_space.NewServerOptions(&server)
	router := generated.HandlerWithOptions(server, options)

	serviceLogger.Info().Msgf("Start SpekaSpace server: %s", address)

	err = http.ListenAndServe(address, router)
	if err != nil {
		serviceLogger.Error().Msg("Error serving http")
	}

	return err
}
