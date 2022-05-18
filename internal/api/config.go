package api

type Config struct {
	Port string `yaml:"port" env:"SERVER_PORT" env-description:"Server port" env-default:"1927"`
}
