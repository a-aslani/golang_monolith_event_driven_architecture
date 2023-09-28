package handlers

import (
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/domain"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
)

func RegisterStoreHandlers(handlers ddd.EventHandler[ddd.AggregateEvent], subscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	subscriber.Subscribe(
		handlers,
		domain.StoreCreatedEvent,
		domain.StoreEditedEvent,
		domain.StoreRemovedEvent,
	)
}
