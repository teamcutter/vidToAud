package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teamcutter/vidToAud/backend/internal/handlers"
)

func SetUpRoutes(s *gin.Engine) {
	s.GET("/:videoID", handlers.Search)
}
