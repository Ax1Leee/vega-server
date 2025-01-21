package app

import (
	"vega-server/pkg/log"
	"vega-server/pkg/server"

	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	server server.Server
	logger *log.Logger
}

type Option func(app *App)

func NewApp(options ...Option) *App {
	app := &App{}
	for _, option := range options {
		option(app)
	}
	return app
}

func WithServer(server server.Server) Option {
	return func(app *App) {
		app.server = server
	}
}

func WithLogger(logger *log.Logger) Option {
	return func(app *App) {
		app.logger = logger
	}
}

func (app *App) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app.logger.Sugar().Info("Starting server...")
	go func() {
		if err := app.server.Start(ctx); err != nil {
			app.logger.Sugar().Errorw("Failed to start server", "error", err)
			cancel()
		}
	}()

	time.Sleep(5 * time.Second)
	select {
	case <-ctx.Done():
		return errors.New("context canceled")
	default:
		app.logger.Sugar().Info("Server started successfully.")
	}

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-signalChannel:
		app.logger.Sugar().Info("Stopping server...")
	}
	if err := app.server.Stop(ctx); err != nil {
		app.logger.Sugar().Errorw("Failed to stop server", "error", err)
		cancel()
	}

	time.Sleep(5 * time.Second)
	select {
	case <-ctx.Done():
		return errors.New("context canceled")
	default:
		app.logger.Sugar().Info("Server stopped successfully.")
	}

	return nil
}
