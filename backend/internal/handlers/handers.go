package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teamcutter/vidToAud/backend/internal/utils"
)

// ! ocRN6Hz3LrI test link !
func Search(c *gin.Context) {
	videoID := c.Param("videoID")
	audioRequired, err := strconv.Atoi(c.Param("audioRequired"))
	if err != nil {
		log.Println(err)
	}
	fileName := utils.DownloadVideo(videoID)
	if audioRequired == 1 {
		utils.ExtractAudio(fileName)
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message" : "Success!",
	})
}
