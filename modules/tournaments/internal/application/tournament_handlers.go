package application

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/domain"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
)

type TournamentHandlers[T ddd.AggregateEvent] struct {
	repo domain.TournamentRepository
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*TournamentHandlers[ddd.AggregateEvent])(nil)

func NewTournamentHandlers(repo domain.TournamentRepository) *TournamentHandlers[ddd.AggregateEvent] {
	return &TournamentHandlers[ddd.AggregateEvent]{
		repo: repo,
	}
}

func (h TournamentHandlers[T]) HandleEvent(ctx context.Context, event T) error {

	switch event.EventName() {
	case domain.TournamentCreatedEvent:
		return h.onTournamentCreated(ctx, event)
	}

	return nil
}

func (h TournamentHandlers[T]) onTournamentCreated(ctx context.Context, event ddd.AggregateEvent) error {

	payload := event.Payload().(*domain.TournamentCreated)

	return h.repo.InsertTournament(ctx, event.AggregateID(), payload.Name.Value, payload.Description.Value, payload.Gamer1ID.Value, payload.Gamer2ID.Value, payload.CreatedAt)
}
