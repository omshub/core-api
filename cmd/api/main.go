package main

import (
	"context"
	"log"
	"net/http"
	"omshub/core-api/internal/api"
	"os/signal"
	"syscall"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {

	var cfg api.Config

	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatalf("[error] config: %s\n", err)
	}

	server := api.NewServer(cfg)

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
