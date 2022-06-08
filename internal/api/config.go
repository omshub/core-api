package api

import "time"

type Config struct {
	Port            string        `env:"PORT" env-description:"Server port" env-default:"1927"`
	NewRelicAPIKey  string        `env:"NEWRELIC_API_KEY" env-description:"API key for NewRelic"`
	NewRelicAppName string        `yaml:"newrelic_app_name" env:"NEWRELIC_APP_NAME" env-description:"App name for NewRelic"`
	DATABASE_URL    string        `env:"DATABASE_URL" env-description:"Database URL" env-default:"host=localhost user=postgres password=postgres dbname=postgres port=5432"`
	ShutdownTimeout time.Duration `yaml:"shutdownTimeout" env:"SHUTDOWN_TIMEOUT" env-description:"Duration for shutdown before force terminate" env-default:"5s"`
}
