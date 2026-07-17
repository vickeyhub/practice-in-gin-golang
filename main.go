package main

import (
	"learn-go/middleware"
	"learn-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Logger())
	// router.Use(middleware.CheckAPIKey())
	routes.RegisterRoutes(router)

	router.Run(":8002")
}
