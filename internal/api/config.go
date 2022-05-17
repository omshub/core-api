package api

type Config struct {
	Port            string `yaml:"port" env:"SERVER_PORT" env-description:"Server port" env-default:"1927"`
	NewRelicAPIKey  string `env:"NEWRELIC_API_KEY" env-description:"API key for NewRelic"`
	NewRelicAppName string `yaml:"newrelic_app_name" env:"NEWRELIC_APP_NAME" env-description:"App name for NewRelic"`
}
