package main

import (
	message_bus "kuchichan/kafka-lp/internal/msg-bus"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const version = "1.0.0"

type application struct {
	env            string
	version        string
	kafkaPublisher *message_bus.KafkaPublisher
}

func main() {
	app := application{
		env:            "development",
		version:        "1.0.0",
		kafkaPublisher: message_bus.InitPublisher(),
	}
	httpRouter := httprouter.New()

	httpRouter.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheck)
	httpRouter.HandlerFunc(http.MethodPost, "/v1/orders", app.createOrder)

	server := http.Server{
		Addr:    ":6000",
		Handler: httpRouter,
	}

	log.Fatal(server.ListenAndServe())
}
