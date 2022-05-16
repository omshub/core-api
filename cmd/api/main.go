package main

import (
	"omshub/core-api/internal/api"
)

func main() {
	server := api.NewServer(api.Config{
		Port: "1927",
	})

	_ = server.Serve()
}
