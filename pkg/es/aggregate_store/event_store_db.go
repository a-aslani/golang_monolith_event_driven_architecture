package aggregate_store

import (
	"context"
	"fmt"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/es"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/registry"
	"github.com/stackus/errors"
	"io"
	"math"
	"time"
)

const (
	count = math.MaxInt64
)

type EventStoreDB struct {
	registry registry.Registry
	db       *esdb.Client
}

type aggregateEvent struct {
	id            string
	name          string
	eventID       string
	eventType     string
	data          []byte
	timestamp     time.Time
	aggregateType ddd.AggregateType
	aggregateID   string
	version       int64
	metadata      []byte
	aggregate     es.EventSourcedAggregate
	payload       ddd.EventPayload
}

var _ es.AggregateStore = (*EventStoreDB)(nil)
var _ ddd.AggregateEvent = (*aggregateEvent)(nil)

func NewEventStoreDB(db *esdb.Client, registry registry.Registry) EventStoreDB {
	return EventStoreDB{
		registry: registry,
		db:       db,
	}
}

func (e EventStoreDB) NewEventFromRecorded(event *esdb.RecordedEvent) (aggregateEvent, error) {

	var payload interface{}
	payload, err := e.registry.Deserialize(event.EventType, event.Data)
	if err != nil {
		return aggregateEvent{}, err
	}

	return aggregateEvent{
		eventID:     event.EventID.String(),
		eventType:   event.EventType,
		data:        event.Data,
		timestamp:   event.CreatedDate,
		aggregateID: event.StreamID,
		version:     int64(event.EventNumber),
		metadata:    event.UserMetadata,
		payload:     payload,
		name:        event.EventType,
	}, nil
}

func (e EventStoreDB) Load(ctx context.Context, aggregate es.EventSourcedAggregate) error {

	stream, err := e.db.ReadStream(ctx, aggregate.ID(), esdb.ReadStreamOptions{}, count)
	if err != nil {
		return errors.Wrap(err, "db.ReadStream")
	}
	defer stream.Close()

	for {
		event, err := stream.Recv()
		if errors.Is(err, esdb.ErrStreamNotFound) {
			return errors.Wrap(err, "stream.Recv")
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return errors.Wrap(err, "stream.Recv")
		}

		esEvent, err := e.NewEventFromRecorded(event.Event)
		if err != nil {
			return errors.Wrap(err, "NewEventFromRecorded(")
		}

		if err = es.LoadEvent(aggregate, esEvent); err != nil {
			return errors.Wrap(err, "LoadEvent")
		}
	}

	return nil
}

func (e EventStoreDB) Save(ctx context.Context, aggregate es.EventSourcedAggregate) error {

	eventsData := make([]esdb.EventData, 0, len(aggregate.Events()))
	for _, event := range aggregate.Events() {

		var dataBytes []byte

		dataBytes, err := e.registry.Serialize(event.EventName(), event.Data())
		if err != nil {
			return err
		}

		eventsData = append(eventsData, esdb.EventData{
			EventType:   event.EventType(),
			ContentType: esdb.JsonContentType,
			Data:        dataBytes,
			Metadata:    event.Metadata(),
		})
	}

	var expectedRevision esdb.ExpectedRevision
	if aggregate.Version() == 0 {
		expectedRevision = esdb.NoStream{}

		appendStream, err := e.db.AppendToStream(
			ctx,
			aggregate.ID(),
			esdb.AppendToStreamOptions{ExpectedRevision: expectedRevision},
			eventsData...,
		)
		if err != nil {
			return errors.Wrap(err, "db.AppendToStream")
		}

		fmt.Printf("(Save) stream: {%+v}\n", appendStream)
		return nil
	}

	readOps := esdb.ReadStreamOptions{Direction: esdb.Backwards, From: esdb.End{}}
	stream, err := e.db.ReadStream(context.Background(), aggregate.ID(), readOps, 1)
	if err != nil {
		return errors.Wrap(err, "db.ReadStream")
	}
	defer stream.Close()

	lastEvent, err := stream.Recv()
	if err != nil {
		return errors.Wrap(err, "stream.Recv")
	}

	expectedRevision = esdb.Revision(lastEvent.OriginalEvent().EventNumber)
	//fmt.Printf("(Save) expectedRevision: {%T}", expectedRevision)

	appendStream, err := e.db.AppendToStream(
		ctx,
		aggregate.ID(),
		esdb.AppendToStreamOptions{ExpectedRevision: expectedRevision},
		eventsData...,
	)
	if err != nil {
		return errors.Wrap(err, "db.AppendToStream")
	}

	fmt.Printf("(Save) stream: {%+v}\n", appendStream)
	//aggregate.ClearUncommittedEvents()
	return nil
}

func (e aggregateEvent) EventName() string {
	return e.AggregateName()
}

func (e aggregateEvent) Payload() ddd.EventPayload {
	return e.payload
}

func (e aggregateEvent) ID() string {
	return e.id
}

func (e aggregateEvent) EventID() string {
	return e.eventID
}

func (e aggregateEvent) EventType() string {
	return e.eventType
}

func (e aggregateEvent) Data() ddd.EventData {
	return e.data
}

func (e aggregateEvent) Timestamp() time.Time {
	return e.timestamp
}

func (e aggregateEvent) AggregateID() string {
	return e.aggregateID
}

func (e aggregateEvent) Metadata() []byte {
	return e.metadata
}

func (e aggregateEvent) AggregateName() string {
	return e.name
}

func (e aggregateEvent) AggregateVersion() int64 {
	return e.version
}

func (e aggregateEvent) Version() int64 {
	return e.version
}
