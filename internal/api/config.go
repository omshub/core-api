package api

import "time"

type Config struct {
	Port            string        `yaml:"port" env:"SERVER_PORT" env-description:"Server port" env-default:"1927"`
	ShutdownTimeout time.Duration `yaml:"shutdownTimeout" env:"SHUTDOWN_TIMEOUT" env-description:"Duration for shutdown before force terminate" env-default:"5s"`
}
