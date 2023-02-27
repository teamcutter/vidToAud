package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/teamcutter/vidToAud/backend/internal/utils"
)

// ! ocRN6Hz3LrI test link !
func GetVideoAndAudio(c *gin.Context) {
	videoID := c.Param("videoID")
	fileName := utils.DownloadVideo(videoID)
	utils.ExtractAudio(fileName)

	c.FileAttachment("/backend/static/downloaded_audio/"+fileName+".mp3", fileName+".mp3")
	/* c.IndentedJSON(http.StatusOK, gin.H{
		"audio": "backend/static/downloaded_audio/" + fileName + ".mp3",
		"video": "backend/static/downloaded_video/" + fileName + ".mp4",
		"filename": fileName,
	}) */
}
