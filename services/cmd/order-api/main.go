package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const version = "1.0.0"

type application struct {
	env     string
	version string
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
