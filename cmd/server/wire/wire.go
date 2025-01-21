//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"vega-server/internal/handler"
	"vega-server/internal/repository"
	"vega-server/internal/server"
	"vega-server/internal/service"
	"vega-server/pkg/app"
	"vega-server/pkg/config"
	"vega-server/pkg/jwt"
	"vega-server/pkg/log"
	"vega-server/pkg/server/http"
)

var RepositorySet = wire.NewSet(
	repository.NewRepository,
	repository.NewDB,
	repository.NewRedis,
	repository.NewUserRepository,
	repository.NewMovieRepository,
	repository.NewReviewRepository,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewMovieService,
	service.NewReviewService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewMovieHandler,
	handler.NewReviewHandler,
)

var ServerSet = wire.NewSet(
	server.NewHTTPServer,
)

func NewApp(httpServer *http.Server, logger *log.Logger) *app.App {
	return app.NewApp(
		app.WithServer(httpServer),
		app.WithLogger(logger),
	)
}

func InitializeApp(conf *config.Config) (*app.App, error) {
	wire.Build(
		RepositorySet,
		ServiceSet,
		HandlerSet,
		ServerSet,
		jwt.NewJWTService,
		log.NewLogger,
		NewApp,
	)
	return &app.App{}, nil
}
