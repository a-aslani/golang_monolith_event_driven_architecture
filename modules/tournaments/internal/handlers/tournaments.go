package handlers

import (
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/domain"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
)

func RegisterTournamentHandlers(handlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(
		handlers,
		domain.TournamentCreatedEvent,
	)
}
