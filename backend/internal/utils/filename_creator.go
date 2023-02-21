package utils

import (
	"fmt"
	"time"
)

func CreateFilename(videoID string, createdAt time.Time) string {
	return fmt.Sprintf("id_%s_%s", videoID, createdAt.Format("2006-01-02-15-04-05"))
}
