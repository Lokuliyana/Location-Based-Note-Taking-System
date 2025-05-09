package routes

import (
	"GeoTagger/controllers"
	"GeoTagger/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleware())
	{
		api.GET("/notes", controllers.GetNotes)    // to implement
		api.POST("/notes", controllers.CreateNote) // to implement
	}
}
