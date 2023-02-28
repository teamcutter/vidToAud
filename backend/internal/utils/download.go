package utils

import (
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/kkdai/youtube/v2"
	"github.com/teamcutter/vidToAud/backend/internal/config"
	"github.com/ztrue/tracerr"
)

func DownloadVideo(videoURL string) string {
	client := youtube.Client{}

	video, err := client.GetVideo(videoURL)
	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
	}
	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
	}
	defer stream.Close()

	filename := CreateFilename(videoURL, video.Title, time.Now())
	file, err := os.Create(config.Conf.VID_PATH + filename + ".mp4")

	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
	}
	log.Println("File installed!")
	return filename
}

func ExtractAudio(filename string) {
	cmd := exec.Command("ffmpeg", "-y", "-i", config.Conf.VID_PATH+filename+".mp4", config.Conf.AUD_PATH+filename+".mp3")
	if err := cmd.Run(); err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
	}
}
