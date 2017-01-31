package main

import (
	"log"
	"net/http"

	"github.com/jvrplmlmn/nginx-requests-stats/handlers"
)

const version = "0.1.0"

func main() {
	log.Println("Starting 'nginx-requests-stats' app...")

	// This endpoint returns a JSON with the version of the application
	http.Handle("/version", handlers.VersionHandler(version))
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
