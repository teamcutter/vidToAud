package utils

import (
	"fmt"
	"regexp"
	"time"
)

func CreateFilename(videoID, videoTitle string, createdAt time.Time) string {
	re := regexp.MustCompile(`[^a-zA-Za-åa-ö-w-я0-9]`)
	title := re.ReplaceAllString(videoTitle, "_")
	return fmt.Sprintf("%s_%s_%s", createdAt.Format("2006-01-02-15-04-05"), title, videoID)
}
