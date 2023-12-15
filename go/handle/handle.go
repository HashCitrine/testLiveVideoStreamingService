package handle

import (
	"errors"
	"hashCitrine/golangHlsServer/service"
	"log"
	"net/http"
	"os"
)

func StreamVideo(h http.Handler) http.HandlerFunc {
	_, err := os.Stat(service.GetOutputFilePath())

	return func(w http.ResponseWriter, r *http.Request) {
		if errors.Is(err, os.ErrExist) {
			log.Fatal("File Not Found")
			service.WaitForPlaylist(service.GetOutputDir(), service.GetOutputFileName())
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

func ConvertVideo(w http.ResponseWriter, r *http.Request) {
	log.Printf("Start ConvertVideo")

	segmentDuration := 1 // duration of each segment in seconds

	if err := service.CreateHLS(service.GetInputFilePath(), service.GetOutputDir(), segmentDuration); err != nil {
		log.Fatalf("Error creating HLS: %v", err)
	}

	log.Println("HLS created successfully")
}
