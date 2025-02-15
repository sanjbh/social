package main

import (
	"log"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	/* msg := struct {
		Status string
	}{
		Status: "ok",
	} */

	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": VERSION,
	}

	if err := writeJSON(w, r, http.StatusOK, data); err != nil {
		log.Printf("error while encoding to JSON: %s\n", err.Error())
	}
}
