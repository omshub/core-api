package api

type Config struct {
	Port            string `yaml:"port" env:"SERVER_PORT" env-description:"Server port" env-default:"1927"`
	NewRelicAPIKey  string `env:"NEWRELIC_API_KEY" env-description:"API key for NewRelic"`
	NewRelicAppName string `yaml:"newrelic_app_name" env:"NEWRELIC_APP_NAME" env-description:"App name for NewRelic"`
	PortDB          string `yaml:"db_port" env:"DB_PORT" env-description:"Database port" env-default:"5432"`
	HostDB          string `yaml:"db_host" env:"DB_HOST" env-description:"Database hostname" env-default:"localhost"`
	NameDB          string `yaml:"db_name" env:"DB_NAME" env-description:"Database name" env-default:"postgres"`
	UserDB          string `yaml:"db_user" env:"DB_USER" env-description:"Database user" env-default:"postgres"`
	PasswordDB      string `yaml:"db_password" env:"DB_PASSWORD" env-description:"Database password" env-default:"postgres"`
}
