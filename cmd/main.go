package main

import (
	"log"
	"logger/internal/handlers"
	"logger/internal/mongo"
	"logger/internal/server"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	err := mongo.ConnectMongoDB()
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/data", handlers.DataHandler)
	r.GET("/data", handlers.GetDataHandler)

	// can be used to start the server

	server.StartServer(r, ":8080")
}
