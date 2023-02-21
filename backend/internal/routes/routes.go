package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teamcutter/vidToAud/backend/internal/config"
	"github.com/teamcutter/vidToAud/backend/internal/handlers"
	"github.com/teamcutter/vidToAud/backend/internal/middlewares"
)

func SetUpRoutes(s *gin.Engine, conf config.Configuration) {
	s.GET("/:videoID/:audioRequired", handlers.Search)
	s.Use(middlewares.Authorization(conf.Token)).POST("/:videoID/:audioRequired", handlers.CleanCache)
}
