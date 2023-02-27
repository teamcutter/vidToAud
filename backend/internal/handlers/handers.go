package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teamcutter/vidToAud/backend/internal/utils"
)

// ! ocRN6Hz3LrI test link !
// ! req == 0 -> video
// ! req == 1 -> audio
// ! req == 2 -> both
func GetVideoAndAudio(c *gin.Context) {
	videoID := c.Param("videoID")
	fileName := utils.DownloadVideo(videoID)
	utils.ExtractAudio(fileName)

	c.IndentedJSON(http.StatusOK, gin.H{
		"audio": "backend/static/downloaded_audio/" + fileName + ".mp3",
		"video": "backend/static/downloaded_video/" + fileName + ".mp4",
	})
}
