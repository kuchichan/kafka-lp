package main

import (
	"kuchichan/kafka-lp/internal/events"
	"kuchichan/kafka-lp/internal/models"
	"kuchichan/kafka-lp/internal/msg-bus"
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
	pb := message_bus.InitPublisher()
	defer pb.Close()

	err := pb.PublishMessage("order-received", event)
	if err != nil {
		panic(err)
	}
}
