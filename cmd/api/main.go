package main

import (
	"log"
	"omshub/core-api/internal/api"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {

	var cfg api.Config

	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatalf("[error] config: %s\n", err)
	}

	server := api.NewServer(cfg)

	_ = server.Serve()
}
