package utils

import (
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/kkdai/youtube/v2"
)

var VID_DIR string = "./downloaded_video/"
var AUD_DIR string = "./downloaded_audio/"

func DownloadVideo(videoID string) string {
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
	defer stream.Close()

	filename := CreateFilename(videoID, video.Title, time.Now())
	file, err := os.Create(VID_DIR + filename + ".mp4")

	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		log.Println(err)
	}
	log.Println("File installed!")
	return filename
}

func ExtractAudio(fileName string) {
	cmd := exec.Command("ffmpeg", "-y", "-i", VID_DIR+fileName+".mp4", AUD_DIR+fileName+".mp3")
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
}
