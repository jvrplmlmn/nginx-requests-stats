package handlers

import (
	"encoding/json"
	"net/http"
)

// VersionResponse ...
type VersionResponse struct {
	Version string `json:"version"`
}

type versionHandler struct {
	version string
}

func (h *versionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := VersionResponse{
		Version: h.version,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// VersionHandler ...
func VersionHandler(version string) http.Handler {
	return &versionHandler{
		version: version,
	}
}
