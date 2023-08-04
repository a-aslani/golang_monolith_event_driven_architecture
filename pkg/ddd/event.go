package ddd

import (
	"time"

	"github.com/google/uuid"
)

type EventPayload interface{}
type EventData interface{}

type Event interface {
	IDer
	EventName() string
	EventType() string
	Data() EventData
	Timestamp() time.Time
	Metadata() Metadata
	Payload() EventPayload
	Version() int64
}

type event struct {
	Entity
	eventType string
	data      EventData
	timestamp time.Time
	metadata  Metadata
	payload   EventPayload
	version   int64
}

var _ Event = (*event)(nil)

func NewEvent(name string, payload EventPayload, options ...EventOption) event {
	return newEvent(name, payload, options...)
}

func newEvent(name string, payload EventPayload, options ...EventOption) event {

	evt := event{
		Entity:    NewEntity(uuid.New().String(), name),
		data:      payload,
		payload:   payload,
		timestamp: time.Now(),
		eventType: name,
		metadata:  make(Metadata),
	}

	for _, option := range options {
		option.configureEvent(&evt)
	}

	return evt
}

func (e event) EventName() string     { return e.name }
func (e event) EventType() string     { return e.eventType }
func (e event) Data() EventData       { return e.data }
func (e event) Timestamp() time.Time  { return e.timestamp }
func (e event) Metadata() Metadata    { return e.metadata }
func (e event) Payload() EventPayload { return e.payload }
func (e event) Version() int64        { return e.version }
