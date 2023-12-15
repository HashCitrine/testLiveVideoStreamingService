package service

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
)

func WaitForPlaylist(dirPath string, fileName string) {
	// 파일 변경 이벤트를 감지할 채널 생성
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// 디렉터리 모니터링을 고루틴으로 실행
	go watchDir(watcher, dirPath)

	// 파일 생성을 기다리는 함수 호출
	waitForFileCreation(watcher, fileName)
}

func watchDir(watcher *fsnotify.Watcher, dirPath string) {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return watcher.Add(path)
	})

	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				log.Printf("File created: %s\n", event.Name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Error:", err)
		}
	}
}

func waitForFileCreation(watcher *fsnotify.Watcher, targetFileName string) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create && filepath.Base(event.Name) == targetFileName {
				log.Printf("Target file %s created!\n", targetFileName)
				return
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Error:", err)
		}
	}
}
