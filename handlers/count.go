package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jvrplmlmn/nginx-requests-stats/stats"
	"github.com/satyrius/gonx"
)

// CountResponse ...
type CountResponse struct {
	Requests int `json:"requests"`
}

type countHandler struct {
	parser  *gonx.Parser
	logFile string
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requests, err := stats.Counter24h(h.parser, h.logFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := CountResponse{
		Requests: requests,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CountHandler ...
func CountHandler(parser *gonx.Parser, logFile string) http.Handler {
	return &countHandler{
		parser:  parser,
		logFile: logFile,
	}
}
