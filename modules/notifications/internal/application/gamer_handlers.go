package application

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/gamerspb"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
	"github.com/rs/zerolog"
)

type GamerHandlers[T ddd.Event] struct {
	logger zerolog.Logger
}

var _ ddd.EventHandler[ddd.Event] = (*GamerHandlers[ddd.Event])(nil)

func NewGamerHandlers(logger zerolog.Logger) GamerHandlers[ddd.Event] {
	return GamerHandlers[ddd.Event]{
		logger: logger,
	}
}

func (h GamerHandlers[T]) HandleEvent(ctx context.Context, event T) error {

	switch event.EventName() {

	case gamerspb.GamerCreatedEvent:
		return h.onGamerCreated(ctx, event)
	}

	return nil
}

func (h GamerHandlers[T]) onGamerCreated(ctx context.Context, event ddd.Event) error {

	payload := event.Payload().(*gamerspb.GamerCreated)

	h.logger.Debug().Msgf(`Send email notification for (ID: %s, FirstName: "%s", LastName: "%s", Email: "%s")`, payload.GetId(), payload.FirstName, payload.LastName, payload.Email)

	return nil
}
