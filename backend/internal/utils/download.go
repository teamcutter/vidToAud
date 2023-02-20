package utils

import (
	"io"
	"log"
	"os"
	"github.com/kkdai/youtube/v2"
)

func Download(videoID string) {
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		log.Println(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		log.Println(err)
	}
	
	/* createdAt := time.Now().Format("01-02-2006-15:04:05")
	if err != nil {
		log.Println(err)
	}
	log.Println(createdAt) */

	file, err := os.Create("./downloaded/" + videoID + ".mp4")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		log.Println(err)
	}
}
