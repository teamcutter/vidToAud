package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teamcutter/vidToAud/backend/internal/utils"
)

// ! ocRN6Hz3LrI test link !
func GetAudioFromVideo(c *gin.Context) {
	videoURL := c.Param("videoURL")
	fileName := utils.DownloadVideo(videoURL)
	utils.ExtractAudio(fileName)
	c.IndentedJSON(http.StatusOK, gin.H{
		"audio":    "/backend/static/downloaded_audio/" + fileName + ".mp3",
		"filename": fileName,
	})
}
