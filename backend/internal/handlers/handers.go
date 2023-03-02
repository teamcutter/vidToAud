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
	c.File(config.Conf.AUD_PATH + fileName)
}
