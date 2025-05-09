package main

import (
	"GeoTagger/config"
	"GeoTagger/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080") // Start server
}
