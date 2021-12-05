package main

import (
	"context"
	"fmt"
	"git.redmadrobot.com/internship/backend/lim-ext/src"
	"git.redmadrobot.com/internship/backend/lim-ext/src/config"
	"git.redmadrobot.com/internship/backend/lim-ext/src/generated"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/db"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/logger"
	"git.redmadrobot.com/internship/backend/lim-ext/src/repository"
	"git.redmadrobot.com/internship/backend/lim-ext/src/repository/ent"
	"git.redmadrobot.com/internship/backend/lim-ext/src/repository/ent/migrate"

	"github.com/rs/zerolog"
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
	journal, err := logger.NewLogger(serviceConfig.Logger)
	if err != nil {
		return err
	}

	// Database
	journal.Info().Msg(serviceConfig.DB.String())

	dbDriver, err := db.GetDriver(serviceConfig.DB, journal)
	if err != nil {
		journal.Error().Err(err).Msg("Failed connect to DB")

		return err
	}

	dbClient, err := speka_space.NewDBClient(dbDriver, journal, serviceConfig.DB.Debug)
	if err != nil {
		journal.Error().Err(err).Msg("Failed get ent client")

		return err
	}

	defer func(dbClient *ent.Client) {
		err := dbClient.Close()
		if err != nil {
			journal.Error().Err(err).Msg("Error while closing DB client")
		}
	}(dbClient)

	makeMigrate(dbClient, context.Background(), journal)

	repo := repository.NewRepository(dbClient, journal)

	server := speka_space.NewServer(repo, journal, serviceConfig)
	options := speka_space.NewServerOptions(&server)
	router := generated.HandlerWithOptions(server, options)

	journal.Info().Msgf("Start SpekaSpace server: %s", address)

	err = http.ListenAndServe(address, router)
	if err != nil {
		journal.Error().Msg("Error serving http")
	}

	return err
}

func makeMigrate(dbClient *ent.Client, ctx context.Context, journal *zerolog.Logger) {
	//Migrate
	if err := dbClient.Schema.WriteTo(ctx, os.Stdout); err != nil {
		journal.Error().Err(err).Msg("failed creating schema resources")
	}
	// Run migration.
	err := dbClient.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if err != nil {
		journal.Error().Err(err).Msg("Failed to make migration")

		return
	}
}
