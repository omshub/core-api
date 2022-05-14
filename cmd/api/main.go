package main

import (
	"github.com/gin-gonic/gin"
	"omshub/core-api/internal/api"
)

func main() {
	router := gin.Default()
	router.GET("/", api.Index)
	router.GET("/ping", api.Ping)

	router.Run(":3000")
}
