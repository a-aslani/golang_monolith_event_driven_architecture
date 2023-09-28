package application

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/domain"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
)

type StoreHandlers[T ddd.AggregateEvent] struct {
	repo domain.StoreRepository
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*StoreHandlers[ddd.AggregateEvent])(nil)

func NewStoreHandlers(repo domain.StoreRepository) *StoreHandlers[ddd.AggregateEvent] {
	return &StoreHandlers[ddd.AggregateEvent]{
		repo: repo,
	}
}

func (h StoreHandlers[T]) HandleEvent(ctx context.Context, event T) error {

	switch event.EventName() {
	case domain.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case domain.StoreEditedEvent:
		return h.onStoreEditedEvent(ctx, event)
	case domain.StoreRemovedEvent:
		return h.onStoreRemovedEvent(ctx, event)
	}

	return nil
}

func (h StoreHandlers[T]) onStoreCreated(ctx context.Context, event ddd.AggregateEvent) error {

	payload := event.Payload().(*domain.StoreCreated)

	return h.repo.Insert(ctx, event.AggregateID(), payload.Name.Value, payload.Amount.Value, payload.Price.Value)
}

func (h StoreHandlers[T]) onStoreEditedEvent(ctx context.Context, event ddd.AggregateEvent) error {

	payload := event.Payload().(*domain.StoreEdited)

	return h.repo.Update(ctx, event.AggregateID(), payload.Name.Value, payload.Amount.Value, payload.Price.Value)
}

func (h StoreHandlers[T]) onStoreRemovedEvent(ctx context.Context, event ddd.AggregateEvent) error {
	return h.repo.Remove(ctx, event.AggregateID())
}
