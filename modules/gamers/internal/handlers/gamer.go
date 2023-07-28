package handlers

import (
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
)

func RegisterGamerHandlers(handlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(domain.GamerCreatedEvent, handlers)
	domainSubscriber.Subscribe(domain.GamerApprovedEvent, handlers)
	domainSubscriber.Subscribe(domain.GamerDisapprovedEvent, handlers)
}
