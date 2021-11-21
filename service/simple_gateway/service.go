package simple_gateway

import (
	"context"
	"os"

	"git.redmadrobot.com/internship/backend/lim-ext/service/simple_gateway/config"

	"github.com/getkin/kin-openapi/openapi3filter"

	"git.redmadrobot.com/internship/backend/lim-ext/service/simple_gateway/generated"
	chmw "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
)

const APIPrefix = "/api/v1"

var _ generated.ServerInterface = (*Server)(nil)

type Server struct {
	cfg    *config.SimpleGWConfig
	logger zerolog.Logger
}

func NewServer(logger *zerolog.Logger, cfg *config.SimpleGWConfig) Server {
	return Server{
		cfg:    cfg,
		logger: *logger,
	}
}

func NewServerOptions(s *Server) generated.ChiServerOptions {
	swagger, err := generated.GetSwagger()
	if err != nil {
		s.logger.Error().Msg("Swagger fails")
		os.Exit(1)
	}
	// this part needs for validation. / prefix added only for /spec route with static spec files
	swagger.Servers = openapi3.Servers{&openapi3.Server{URL: "/api/v1"}, &openapi3.Server{URL: "/"}}

	r := chi.NewRouter()

	validatorOptions := &chmw.Options{}
	validatorOptions.Options.AuthenticationFunc = func(c context.Context, input *openapi3filter.AuthenticationInput) error {
		// При указании блока security в методах openapi, функция AuthenticationFunc начинает срабатывать раньше, Чем инициализация middleware.
		// Использовать здесь наш middleware нельзя,  ак как все middleware инициализируются после routers, работаем только с нашим middleware авторизации
		return nil
	}

	r.Use(httplog.RequestLogger(s.logger))
	r.Use(chmw.OapiRequestValidatorWithOptions(swagger, validatorOptions))
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	// Publish redoc specification (all from api dir) path actual for current docker build, fix for local
	return generated.ChiServerOptions{
		BaseURL:    APIPrefix,
		BaseRouter: r,
	}
}
