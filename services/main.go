package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const version = "1.0.0"

type application struct {
	env     string
	version string
}

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

func main() {
	app := application{env: "development", version: "1.0.0"}
	httpRouter := httprouter.New()
	httpRouter.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheck)

	server := http.Server{
		Addr:    ":6000",
		Handler: httpRouter,
	}

	log.Fatal(server.ListenAndServe())
}
