package routes

import (
	"logger/internal/handlers"
	"logger/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func LogRoutes(r *gin.Engine) {
	// r.POST("/data", handlers.DataHandler)
	// r.GET("/data", handlers.GetDataHandler)

	data := r.Group("/data")
	data.Use(middlewares.JWTAuthMiddleware)
	{
		data.POST("", handlers.DataHandler)
		data.GET("", handlers.GetDataHandler)
	}

}
