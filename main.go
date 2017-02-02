package main

import (
	"flag"
	"log"
	"net/http"

	h "github.com/jvrplmlmn/nginx-requests-stats/handlers"
	"github.com/satyrius/gonx"
)

const version = "0.1.0"

var format string
var logFile string

func init() {
	flag.StringVar(&format, "format", `$remote_addr - $remote_user [$time_local] "$request" $status $body_bytes_sent "$http_referer" "$http_user_agent" $upstream_addr $upstream_cache_status`, "Log format")
	flag.StringVar(&logFile, "log", "/var/log/nginx/access.log", "Log file name to read.")
}

func main() {
	// Parse the command-line flags
	flag.Parse()

	// Always log when the application starts
	log.Println("Starting 'nginx-requests-stats' app...")

	// Create a parser based on a given format
	parser := gonx.NewParser(format)

	// This endpoint returns a JSON with the version of the application
	h.HandleWLogger("version", h.VersionHandler(version))
	// This endpoint returns a JSON with the number of requests in the last 24h
	h.HandleWLogger("/count", h.CountHandler(parser, logFile))
	// Serve the endpoints
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
