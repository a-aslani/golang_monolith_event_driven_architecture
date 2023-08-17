package commands

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/domain"
	"time"
)

type CreateTournament struct {
	ID          string
	Name        string
	Description string
	Gamer1ID    string
	Gamer2ID    string
	CreatedAt   time.Time
}

type CreateTournamentHandler struct {
	eventStore domain.TournamentEventStore
}

func NewCreateTournamentHandler(eventStore domain.TournamentEventStore) CreateTournamentHandler {
	return CreateTournamentHandler{eventStore: eventStore}
}

func (h CreateTournamentHandler) CreateTournament(ctx context.Context, cmd CreateTournament) error {

	tournament, err := domain.CreateTournament(cmd.ID, cmd.Name, cmd.Description, cmd.Gamer1ID, cmd.Gamer2ID, cmd.CreatedAt)
	if err != nil {
		return err
	}

	return h.eventStore.Save(ctx, tournament)
}
