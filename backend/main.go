package main

import (
	"github.com/teamcutter/vidToAud/backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	routes.SetUpRoutes(server)
	server.Run(":8080")
}
