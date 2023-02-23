package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teamcutter/vidToAud/backend/internal/utils"
	"github.com/ztrue/tracerr"
)

// ! ocRN6Hz3LrI test link !
// ! req == 0 -> video
// ! req == 1 -> audio
// ! req == 2 -> both
func Search(c *gin.Context) {
	videoID := c.Param("videoID")
	requirements, err := strconv.Atoi(c.Param("requirements"))
	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
	}
	fileName := utils.DownloadVideo(videoID)

	switch requirements {
	case 0:
		c.IndentedJSON(http.StatusOK, gin.H{
			"video": "backend/static/downloaded_video/" + fileName + ".mp4",
		})
		return
	case 1:
		utils.ExtractAudio(fileName)
		c.IndentedJSON(http.StatusOK, gin.H{
			"audio": "backend/static/downloaded_audio/" + fileName + ".mp3",
		})
		return
	case 2:
		utils.ExtractAudio(fileName)
		c.IndentedJSON(http.StatusOK, gin.H{
			"audio": "backend/static/downloaded_audio/" + fileName + ".mp3",
			"video": "backend/static/downloaded_video/" + fileName + ".mp4",
		})
		return
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"message": "bad request",
	})
}
