package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Server struct {
	config     Config
	router     *gin.Engine
	httpServer *http.Server
}

func NewServer(config Config) *Server {
	router := gin.Default()

	// TODO: improve error handling and error logging
	if config.NewRelicAPIKey != "" && config.NewRelicAppName != "" {
		app, err := newRelicApplication(config.NewRelicAppName, config.NewRelicAPIKey)
		if err != nil {
			log.Printf("error configuring NewRelic: %v", err)
		} else {
			router.Use(nrgin.Middleware(app))
		}
	} else {
		log.Printf("NewRelic configuration not found. Code will not be instrumented.")
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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

func newRelicApplication(appName string, apiKey string) (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(apiKey),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
}
