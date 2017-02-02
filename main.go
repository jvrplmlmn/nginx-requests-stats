package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	h "github.com/jvrplmlmn/nginx-requests-stats/handlers"
	"github.com/kelseyhightower/envconfig"
	"github.com/satyrius/gonx"
)

const version = "0.3.0"

var format string
var logFile string

// Configuration holds the values used to configure the application
// In this case we're only using it for the HTTP Server
type Configuration struct {
	HTTPAddr string `default:"localhost"`
	HTTPPort int    `default:"8080"`
}

func init() {
	flag.StringVar(&format, "format", `$remote_addr - $remote_user [$time_local] "$request" $status $body_bytes_sent "$http_referer" "$http_user_agent" $upstream_addr $upstream_cache_status`, "Log format")
	flag.StringVar(&logFile, "log", "/var/log/nginx/access.log", "Log file name to read.")
}

func main() {
	// Parse the command-line flags
	flag.Parse()

	// Read and process the configuration from environment variables
	var cnfg Configuration

	err := envconfig.Process("nginx_stats", &cnfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Based on the configuration, compose the address
	// where we need to bind the HTTP Server
	bind := fmt.Sprintf("%v:%v", cnfg.HTTPAddr, cnfg.HTTPPort)

	// Always log when the application starts
	log.Printf("Starting 'nginx-requests-stats' app (version: %s)\n", version)

	// Create a parser based on a given format
	parser := gonx.NewParser(format)

	// This endpoint returns a JSON with the version of the application
	h.HandleWLogger("version", h.VersionHandler(version))
	// This endpoint returns a JSON with the number of requests in the last 24h
	h.HandleWLogger("/count", h.CountHandler(parser, logFile))
	// Serve the endpoints
	log.Fatal(http.ListenAndServe(bind, nil))

}
