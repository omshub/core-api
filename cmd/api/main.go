package main

import (
	"omshub/core-api/internal/api"
	"os"
)

func main() {
	server := api.NewServer(api.Config{
		Port:            "1927",
		NewRelicAPIKey:  os.Getenv("NEWRELIC_API_KEY"),
		NewRelicAppName: os.Getenv("NEWRELIC_APP_NAME"),
	})

	_ = server.Serve()
}
