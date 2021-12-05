package speka_space

import (
	"context"
	entSql "entgo.io/ent/dialect/sql"
	"git.redmadrobot.com/internship/backend/lim-ext/src/config"
	gen "git.redmadrobot.com/internship/backend/lim-ext/src/generated"
	"git.redmadrobot.com/internship/backend/lim-ext/src/repository/ent"
	entMigrate "git.redmadrobot.com/internship/backend/lim-ext/src/repository/ent/migrate"
	"os"

	"github.com/getkin/kin-openapi/openapi3filter"

	repo "git.redmadrobot.com/internship/backend/lim-ext/src/repository"

	chimiddleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
)

const APIPrefix = "/api/v1"

var _ gen.ServerInterface = (*Server)(nil)

type Server struct {
	cfg    *config.SpekaSpaceConfig
	logger zerolog.Logger
	repo   repo.Repository
}

func NewServer(repository repo.Repository, logger *zerolog.Logger, cfg *config.SpekaSpaceConfig) Server {
	return Server{
		repo:   repository,
		cfg:    cfg,
		logger: *logger,
	}
}

func NewServerOptions(s *Server) gen.ChiServerOptions {
	swagger, err := gen.GetSwagger()
	if err != nil {
		s.logger.Error().Msg("Swagger fails")
		os.Exit(1)
	}
	swagger.Servers = openapi3.Servers{&openapi3.Server{URL: "/api/v1"}, &openapi3.Server{URL: "/"}}

	r := chi.NewRouter()

	validatorOptions := &chimiddleware.Options{}
	validatorOptions.Options.AuthenticationFunc = func(c context.Context, input *openapi3filter.AuthenticationInput) error {
		return nil
	}

	r.Use(httplog.RequestLogger(s.logger))
	r.Use(chimiddleware.OapiRequestValidatorWithOptions(swagger, validatorOptions))
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	return gen.ChiServerOptions{
		BaseURL:    APIPrefix,
		BaseRouter: r,
	}
}

func NewDBClient(drv *entSql.Driver, log *zerolog.Logger, debug bool) (client *ent.Client, err error) {
	client = ent.NewClient(ent.Driver(drv))
	if debug {
		err = client.Debug().Schema.Create(
			context.Background(),
			entMigrate.WithDropColumn(true),
		)
	} else {
		err = client.Schema.Create(
			context.Background(),
			entMigrate.WithDropColumn(true),
		)
	}
	if err != nil {
		log.Error().Err(err).Msg(err.Error())

		return nil, err
	}

	return client, nil
}
