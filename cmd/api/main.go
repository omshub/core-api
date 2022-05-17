package main

import (
	"log"

	"omshub/core-api/internal/api"

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

	server := api.NewServer(cfg, serverDeps)

	_ = server.Serve()
}

func newRelicApplication(appName string, apiKey string) (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(apiKey),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
}
