package domain

import (
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/domain/value_objects"
	"time"
)

const (
	TournamentCreatedEvent = "V1.Tournaments.TournamentCreated"
)

type TournamentCreated struct {
	Name        value_objects.TournamentName
	Description value_objects.TournamentDescription
	Gamer1ID    value_objects.GamerID
	Gamer2ID    value_objects.GamerID
	CreatedAt   time.Time
}

func (TournamentCreated) Key() string { return TournamentCreatedEvent }
