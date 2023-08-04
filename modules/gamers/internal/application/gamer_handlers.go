package application

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
)

type GamerHandlers[T ddd.AggregateEvent] struct {
	repo domain.GamerRepository
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*GamerHandlers[ddd.AggregateEvent])(nil)

func NewGamerHandlers(repo domain.GamerRepository) *GamerHandlers[ddd.AggregateEvent] {
	return &GamerHandlers[ddd.AggregateEvent]{
		repo: repo,
	}
}

func (h GamerHandlers[T]) HandleEvent(ctx context.Context, event T) error {

	switch event.EventName() {
	case domain.GamerCreatedEvent:
		return h.onGamerCreated(ctx, event)
	case domain.GamerApprovedEvent:
		return h.onGamerApproved(ctx, event)
	case domain.GamerDisapprovedEvent:
		return h.onGamerDisapproved(ctx, event)
	}

	return nil
}

func (h GamerHandlers[T]) onGamerDisapproved(ctx context.Context, event ddd.AggregateEvent) error {
	return h.repo.ChangeGamerState(ctx, event.AggregateID(), false)
}

func (h GamerHandlers[T]) onGamerApproved(ctx context.Context, event ddd.AggregateEvent) error {
	return h.repo.ChangeGamerState(ctx, event.AggregateID(), true)
}

func (h GamerHandlers[T]) onGamerCreated(ctx context.Context, event ddd.AggregateEvent) error {

	payload := event.Payload().(*domain.GamerCreated)

	return h.repo.CreateGamer(ctx, event.AggregateID(), payload.FirstName, payload.LastName, payload.Email, payload.Password, payload.IsApproved)
}
