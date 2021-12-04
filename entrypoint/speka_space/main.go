package main

import (
	"fmt"
	"git.redmadrobot.com/internship/backend/lim-ext/service/speka_space"
	"net/http"
	"os"

	"git.redmadrobot.com/internship/backend/lim-ext/pkg/logger"

	"git.redmadrobot.com/internship/backend/lim-ext/service/speka_space/config"
	"github.com/rs/zerolog/log"

	"git.redmadrobot.com/internship/backend/lim-ext/service/speka_space/generated"
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
