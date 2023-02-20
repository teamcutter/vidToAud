package handlers

import (
	"log"

	"github.com/teamcutter/vidToAud/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

// ! ocRN6Hz3LrI test link !
func Search(c *gin.Context) {
	videoID := c.Param("videoID")
	go utils.Download(videoID)
	log.Println(videoID)
}
