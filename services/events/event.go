package events

import (
	"time"

	"github.com/google/uuid"
)

type Event interface {
	ID() uuid.UUID
	Name() string
	Timestamp() time.Time
	Body() any
}

type BaseEvent struct {
	EventID        uuid.UUID
	EventTimestamp time.Time
}
