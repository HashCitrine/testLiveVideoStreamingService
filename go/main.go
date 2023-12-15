package main

import (
	"fmt"
	"github.com/HashCitrine/testHls/handle"
	"github.com/HashCitrine/testHls/service"
	"log"
	"net/http"
)

func main() {
	const port = 8080

	http.Handle("/", handle.StreamVideo(http.FileServer(http.Dir(service.GetOutputDir()))))
	http.HandleFunc("/convert", handle.ConvertVideo)

	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port: %v\n", service.GetOutputDir(), port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
