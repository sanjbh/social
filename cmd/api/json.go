package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, r *http.Request, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	return json.NewDecoder(r.Body).Decode(data)
}
