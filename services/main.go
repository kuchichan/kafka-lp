package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/julienschmidt/httprouter"
)

const version = "1.0.0"

type application struct {
	env     string
	version string
}

type OrderReceived struct {
	ID        string `json:"ID"`
	Timestamp int    `json:"timestamp"`
	Name      string `json:"name"`
}

func publishMessage(p *kafka.Producer, topic string, message any) error {
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	js, err := json.Marshal(message)
	if err != nil {
		return err
	}

	p.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          js,
		},
		nil,
	)

	return nil
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
	orderReceived := OrderReceived{
		ID:        "25e85f0f-936d-4edc-9992-ad33611d80fb",
		Name:      "OrderReceived",
		Timestamp: 1698863768,
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})

	if err != nil {
		panic(err)
	}

	defer p.Close()

	publishMessage(p, "order-received", orderReceived)

	app := application{env: "development", version: "1.0.0"}
	httpRouter := httprouter.New()
	httpRouter.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheck)

	server := http.Server{
		Addr:    ":6000",
		Handler: httpRouter,
	}

	log.Fatal(server.ListenAndServe())
}
