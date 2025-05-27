package main

import (
	"github.com/gin-gonic/gin"
	"hnex.com/internal/config"
)

func main() {
	router := gin.Default()

	config.SetupRouter(router)

	router.Run()
}
