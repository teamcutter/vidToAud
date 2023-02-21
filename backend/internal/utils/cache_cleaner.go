package utils

import (
	"log"
	"os"
	"path"
)

func Clean(dirs ...string) {
	if len(dirs) == 1 {
		files, err := os.ReadDir(dirs[0])
		if err != nil {
			log.Println(err)
		}
		for _, file := range files {
			os.RemoveAll(path.Join(dirs[0], file.Name()))
		}
	} else {
		go Clean(dirs[0])
		go Clean(dirs[1])
	}
}
