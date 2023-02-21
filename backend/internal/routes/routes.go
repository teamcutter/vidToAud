package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teamcutter/vidToAud/backend/internal/config"
	"github.com/teamcutter/vidToAud/backend/internal/handlers"
)

func SetUpRoutes(s *gin.Engine, conf config.Configuration) {
	s.GET("/:videoID/:audioRequired", handlers.Search)
}
