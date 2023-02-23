package main

import (
	"github.com/teamcutter/vidToAud/backend/internal/config"
	"github.com/teamcutter/vidToAud/backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.GetConfig()
	server := gin.Default()
	routes.SetUpRoutes(server, config.Conf)
	server.Run(":8080")
}
