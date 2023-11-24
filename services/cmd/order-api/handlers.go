package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "OK",
		"environment": app.env,
		"version":     app.version,
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Json could not be processed", http.StatusInternalServerError)
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
