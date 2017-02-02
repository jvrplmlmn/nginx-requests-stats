package handlers

import (
	"log"
	"net/http"
	"time"
)

func HandleWLogger(pattern string, handler http.Handler) {
	http.Handle(pattern, Logger(handler))
}

// Logger ...
func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
