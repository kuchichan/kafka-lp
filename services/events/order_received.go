package events

import (
	"kuchichan/kafka-lp/models"
	"time"

	"github.com/google/uuid"
)

type OrderReceived struct {
	BaseEvent
	EventBody models.Order
}

func (order OrderReceived) ID() uuid.UUID {
	return order.EventID
}

func (order OrderReceived) Name() string {
	return "OrderReceived"
}

func (order OrderReceived) Timestamp() time.Time {
	return order.EventTimestamp
}

func (order OrderReceived) Body() models.Order {
	return order.EventBody
}
