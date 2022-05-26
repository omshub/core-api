package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Server struct {
	config     Config
	router     *gin.Engine
	httpServer *http.Server
}

// Dependencies provides a crude means of dependency injection.
type Dependencies struct {
	NewRelicApp *newrelic.Application // All operations on this are nil-safe. No mocks required for testing.
}

func NewServer(config Config, deps Dependencies) *Server {
	router := gin.Default()

	if deps.NewRelicApp != nil {
		router.Use(nrgin.Middleware(deps.NewRelicApp))
	}

	httpServer := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	server := &Server{
		config:     config,
		router:     router,
		httpServer: httpServer,
	}

	router.GET("/", server.Index)
	router.GET("/ping", server.Ping)

	return server
}

func (s *Server) Serve() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.config.ShutdownTimeout)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}

func (s *Server) Index(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (s *Server) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
