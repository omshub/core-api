package api

import "time"

type Config struct {
	Port             string
	ShutdownDuration time.Duration
}

// TODO: use a library to manage sensible defaults via struct tags
func DefaultConfig() Config {
	return Config{
		Port:             "1927",
		ShutdownDuration: time.Second * 5,
	}
}
