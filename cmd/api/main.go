package main

import (
	"context"
	"log"
	"net/http"
	"omshub/core-api/internal/api"
	"os/signal"
	"syscall"
)

func main() {
	// TODO: read from file or env vars
	config := api.DefaultConfig()
	server := api.NewServer(config)

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
