package common

import (
	"kuchichan/kafka-lp/internal/events"
	"kuchichan/kafka-lp/internal/models"
	"time"

	"github.com/google/uuid"
)

func OrderToReceivedEvent(order models.Order) events.OrderReceived {
	return events.OrderReceived{
		EventBody: order,
		BaseEvent: events.BaseEvent{
			EventID:        uuid.New(),
			EventTimestamp: time.Now(),
		},
	}
}
