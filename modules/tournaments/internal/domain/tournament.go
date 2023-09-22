package domain

import (
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/domain/value_objects"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/es"
	"github.com/stackus/errors"
	"time"
)

const TournamentAggregate = "tournaments.Tournament"

type Tournament struct {
	es.Aggregate
	name        value_objects.TournamentName
	description value_objects.TournamentDescription
	gamer1ID    value_objects.GamerID
	gamer2ID    value_objects.GamerID
	createdAt   time.Time
	finishedAt  time.Time
}

var _ interface {
	es.EventApplier
	es.Snapshotter
} = (*Tournament)(nil)

func (Tournament) Key() string { return TournamentAggregate }

func NewTournament(id string) *Tournament {
	return &Tournament{
		Aggregate: es.NewAggregate(id, TournamentAggregate),
	}
}

func CreateTournament(id, name, description, gamer1ID, gamer2ID string, createdAt time.Time) (tournament *Tournament, err error) {

	tournament = NewTournament(id)

	tournament.name, err = value_objects.NewTournamentName(name)
	if err != nil {
		return nil, err
	}

	tournament.description, err = value_objects.NewTournamentDescription(description)
	if err != nil {
		return nil, err
	}

	tournament.gamer1ID, err = value_objects.NewGamerID(gamer1ID)
	if err != nil {
		return nil, err
	}

	tournament.gamer2ID, err = value_objects.NewGamerID(gamer2ID)
	if err != nil {
		return nil, err
	}

	tournament.createdAt = createdAt

	tournament.AddEvent(TournamentCreatedEvent, &TournamentCreated{
		Name:        tournament.name,
		Description: tournament.description,
		Gamer1ID:    tournament.gamer1ID,
		Gamer2ID:    tournament.gamer2ID,
		CreatedAt:   tournament.createdAt,
	})

	return tournament, nil
}

func (a *Tournament) ApplyEvent(event ddd.Event) error {

	switch payload := event.Payload().(type) {
	case *TournamentCreated:

		a.name = payload.Name
		a.description = payload.Description
		a.gamer1ID = payload.Gamer1ID
		a.gamer2ID = payload.Gamer2ID
		a.createdAt = payload.CreatedAt

	default:
		return errors.ErrInternal.Msgf("%T received the event %s with unexpected payload %T", a, event.EventName(), payload)
	}

	return nil
}

func (a *Tournament) ApplySnapshot(snapshot es.Snapshot) (err error) {

	switch ss := snapshot.(type) {
	case *TournamentV1:

		a.name, err = value_objects.NewTournamentName(ss.Name)
		if err != nil {
			return err
		}

		a.description, err = value_objects.NewTournamentDescription(ss.Description)
		if err != nil {
			return err
		}

		a.gamer1ID, err = value_objects.NewGamerID(ss.Gamer1ID)
		if err != nil {
			return err
		}

		a.gamer2ID, err = value_objects.NewGamerID(ss.Gamer2ID)
		if err != nil {
			return err
		}

		a.createdAt = ss.CreatedAt

	default:
		return errors.ErrInternal.Msgf("%T received the unexpected snapshot %T", a, snapshot)
	}
	return nil
}

func (a *Tournament) ToSnapshot() es.Snapshot {
	return TournamentV1{
		Name:        a.name.Value,
		Description: a.description.Value,
		Gamer1ID:    a.gamer1ID.Value,
		Gamer2ID:    a.gamer2ID.Value,
		CreatedAt:   a.createdAt,
	}
}
