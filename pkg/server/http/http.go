package http

import (
	"github.com/gin-gonic/gin"
	"vega-server/pkg/config"

	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	// Embedding
	*gin.Engine
	httpServer *http.Server
	host       string
	port       int
}

type Option func(server *Server)

func NewServer(config *config.Config) *Server {
	server := &Server{
		Engine:     gin.Default(),
		httpServer: &http.Server{},
		host:       config.GetString("http.host"),
		port:       config.GetInt("http.port"),
	}
	server.httpServer.Addr = fmt.Sprintf("%s:%d", server.host, server.port)
	server.httpServer.Handler = server
	return server
}

func (server *Server) Start(ctx context.Context) error {
	if err := server.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return errors.New("failed to start http server")
	}

	return nil
}

func (server *Server) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.httpServer.Shutdown(ctx); err != nil {
		return errors.New("failed to stop http server")
	}

	return nil
}
