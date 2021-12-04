package speka_space

import (
	"context"
	"git.redmadrobot.com/internship/backend/lim-ext/src/config"
	generated2 "git.redmadrobot.com/internship/backend/lim-ext/src/generated"
	"os"

	"github.com/getkin/kin-openapi/openapi3filter"

	chmw "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
)

const APIPrefix = "/api/v1"

var _ generated2.ServerInterface = (*Server)(nil)

type Server struct {
	cfg    *config.SpekaSpaceConfig
	logger zerolog.Logger
}

func NewServer(logger *zerolog.Logger, cfg *config.SpekaSpaceConfig) Server {
	return Server{
		cfg:    cfg,
		logger: *logger,
	}
}

func NewServerOptions(s *Server) generated2.ChiServerOptions {
	swagger, err := generated2.GetSwagger()
	if err != nil {
		s.logger.Error().Msg("Swagger fails")
		os.Exit(1)
	}
	swagger.Servers = openapi3.Servers{&openapi3.Server{URL: "/api/v1"}, &openapi3.Server{URL: "/"}}

	r := chi.NewRouter()

	validatorOptions := &chmw.Options{}
	validatorOptions.Options.AuthenticationFunc = func(c context.Context, input *openapi3filter.AuthenticationInput) error {
		return nil
	}

	r.Use(httplog.RequestLogger(s.logger))
	r.Use(chmw.OapiRequestValidatorWithOptions(swagger, validatorOptions))
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	return generated2.ChiServerOptions{
		BaseURL:    APIPrefix,
		BaseRouter: r,
	}
}
