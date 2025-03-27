package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

// StartServer starts the Gin server on the specified port
func StartServer(r *gin.Engine, port string) {
	log.Printf("Server is running on port %s", port)
	log.Fatal(r.Run(port))
}
