package main

import (
	"log"

	"github.com/teamcutter/vidToAud/backend/internal/config"
	"github.com/teamcutter/vidToAud/backend/internal/middlewares"
	"github.com/teamcutter/vidToAud/backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.GetConfig()
	log.Println("Initialize config...")
}

func main() {
	server := gin.Default()
	server.Use(middlewares.CorsMiddleware())
	routes.SetUpRoutes(server, config.Conf)
	server.Run(":" + config.Conf.DEBUG_PORT)
}
