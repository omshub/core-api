package main

import (
	"omshub/core-api/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", api.Index)
	router.GET("/ping", api.Ping)

	_ = router.Run(":3000")
}
