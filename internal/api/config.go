package api

import "time"

type Config struct {
	Port            string        `yaml:"port" env:"SERVER_PORT" env-description:"Server port" env-default:"1927"`
	ShutdownTimeout time.Duration `yaml:"shutdownTimeout" env:"SHUTDOWN_TIMEOUT" env-description:"Duration for shutdown before force terminate" env-default:"5s"`
	NewRelicAPIKey  string        `env:"NEWRELIC_API_KEY" env-description:"API key for NewRelic"`
	NewRelicAppName string        `yaml:"newrelic_app_name" env:"NEWRELIC_APP_NAME" env-description:"App name for NewRelic"`
}
