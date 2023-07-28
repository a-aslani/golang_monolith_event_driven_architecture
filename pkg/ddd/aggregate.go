package ddd

import "encoding/json"

// AggregateType type of the Aggregate
type AggregateType string

type (
	AggregateNamer interface {
		AggregateName() string
	}

	Eventer interface {
		AddEvent(string, interface{}, ...EventOption)
		Events() []AggregateEvent
		ClearEvents()
	}

	Aggregate struct {
		Entity
		events []AggregateEvent
	}

	AggregateEvent interface {
		Event
		AggregateName() string
		AggregateID() string
		AggregateVersion() int64
	}

	aggregateEvent struct {
		event
	}
)

var _ interface {
	AggregateNamer
	Eventer
} = (*Aggregate)(nil)

func NewAggregate(id, name string) Aggregate {
	return Aggregate{
		Entity: NewEntity(id, name),
		events: make([]AggregateEvent, 0),
	}
}

func (a Aggregate) AggregateName() string    { return a.name }
func (a Aggregate) Events() []AggregateEvent { return a.events }
func (a *Aggregate) ClearEvents()            { a.events = []AggregateEvent{} }

func (a *Aggregate) AddEvent(name string, data interface{}, options ...EventOption) {
	options = append(
		options,
		Metadata{
			AggregateName: a.name,
			AggregateID:   a.id,
			//AggregateVersion: version,
		},
	)
	a.events = append(
		a.events,
		aggregateEvent{
			event: newEvent(name, data, options...),
		},
	)
}

func (a *Aggregate) setEvents(events []AggregateEvent) { a.events = events }

func (e aggregateEvent) AggregateName() string {
	var metadata Metadata
	_ = json.Unmarshal(e.metadata, &metadata)
	return metadata.AggregateName
}
func (e aggregateEvent) AggregateID() string {
	var metadata Metadata
	_ = json.Unmarshal(e.metadata, &metadata)
	return metadata.AggregateID
}
func (e aggregateEvent) AggregateVersion() int64 {
	var metadata Metadata
	_ = json.Unmarshal(e.metadata, &metadata)
	return metadata.AggregateVersion
}
