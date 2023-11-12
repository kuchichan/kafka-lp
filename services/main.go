package main

import "kuchichan/kafka-lp/publisher"

type OrderReceived struct {
	ID        string `json:"ID"`
	Timestamp int    `json:"timestamp"`
	Name      string `json:"name"`
}

func main() {
	orderReceived := OrderReceived{
		ID:        "25e85f0f-936d-4edc-9992-ad33611d80fb",
		Name:      "OrderReceived",
		Timestamp: 1698863768,
	}
	pb := publisher.InitPublisher()
	defer pb.Close()

	err := publisher.PublishMessage(pb, "order-received", orderReceived)
	if err != nil {
		panic(err)
	}
}
