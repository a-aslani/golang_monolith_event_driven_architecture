package application

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/gamerspb"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/am"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
)

type IntegrationEventHandlers[T ddd.AggregateEvent] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*IntegrationEventHandlers[ddd.AggregateEvent])(nil)

func NewIntegrationEventHandlers(publisher am.MessagePublisher[ddd.Event]) *IntegrationEventHandlers[ddd.AggregateEvent] {
	return &IntegrationEventHandlers[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func (h IntegrationEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.GamerCreatedEvent:
		return h.onGamerCreated(ctx, event)
	}
	return nil
}

func (h IntegrationEventHandlers[T]) onGamerCreated(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.GamerCreated)
	return h.publisher.Publish(
		ctx,
		gamerspb.GamerAggregateChannel,
		ddd.NewEvent(
			gamerspb.GamerCreatedEvent,
			&gamerspb.GamerCreated{
				Id:        event.ID(),
				FirstName: payload.FullName.FistName,
				LastName:  payload.FullName.LastName,
				Email:     payload.Email.Value,
			},
		),
	)
}
