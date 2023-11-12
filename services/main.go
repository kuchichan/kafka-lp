package main

import (
	"kuchichan/kafka-lp/events"
	"kuchichan/kafka-lp/models"
	"kuchichan/kafka-lp/publisher"
	"time"

	"github.com/google/uuid"
)

type OrderReceived struct {
	ID        string `json:"ID"`
	Timestamp int    `json:"timestamp"`
	Name      string `json:"name"`
}

func main() {
	order := models.Order{
		ID: uuid.New(),
	}
	event := events.OrderReceived{
		BaseEvent: events.BaseEvent{
			EventID:        uuid.New(),
			EventTimestamp: time.Now(),
		},
		EventBody: order,
	}
	pb := publisher.InitPublisher()
	defer pb.Close()

	err := publisher.PublishMessage(pb, "order-received", event)
	if err != nil {
		panic(err)
	}
}
