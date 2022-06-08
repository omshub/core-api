package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"omshub/core-api/internal/api"
	"omshub/core-api/internal/api/db"
	"os/signal"
	"syscall"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	var cfg api.Config

	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatalf("[error] config: %s\n", err)
	}

	serverDeps := api.Dependencies{}

	if cfg.NewRelicAPIKey == "" || cfg.NewRelicAppName == "" {
		log.Println("[info] NewRelic configuration not provided. Continuing in the dark...")
	}
	if app, err := newRelicApplication(cfg.NewRelicAppName, cfg.NewRelicAPIKey); err != nil {
		log.Printf("[warn] NewRelic could not be configured: %s\n", err)
	} else {
		serverDeps.NewRelicApp = app
	}

	var dsn string
	if strings.HasPrefix(cfg.DATABASE_URL, "host=localhost") {
		dsn = fmt.Sprintf("%s sslmode=disable", cfg.DATABASE_URL)
	} else {
		dsn = fmt.Sprintf("%s?sslmode=require", cfg.DATABASE_URL)
	}

	if db, err := db.NewDB(dsn); err != nil {
		log.Printf("[warn] DB auto migration failed: %s\n", err)
	} else {
		serverDeps.DB = db
	}

	server := api.NewServer(cfg, serverDeps)

	go func() {
		if err := server.Serve(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[error] serve: %s\n", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	stop()

	log.Println("[info] shutting down")

	if err := server.Shutdown(); err != nil {
		log.Fatalf("[error] shutdown: %s\n", err)
	}
}

func newRelicApplication(appName string, apiKey string) (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(apiKey),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
}
