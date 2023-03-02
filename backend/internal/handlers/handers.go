package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/teamcutter/vidToAud/backend/internal/config"
	"github.com/teamcutter/vidToAud/backend/internal/utils"
)

// ! ocRN6Hz3LrI test link !
func GetAudioFromVideo(c *gin.Context) {
	videoURL := c.Param("videoURL")
	fileName := utils.DownloadVideo(videoURL)
	utils.ExtractAudio(fileName)
	// Set the Content-Disposition header to attachment
	c.Header("Content-Disposition", "attachment; filename="+fileName+".mp3")
	// Set the Content-Type header to audio/mpeg
	c.Header("Content-Type", "audio/mpeg")
	c.File(config.Conf.AUD_PATH + fileName + ".mp3")
}
