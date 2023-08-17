package domain

import "context"

type TournamentEventStore interface {
	Save(ctx context.Context, tournament *Tournament) error
	Load(ctx context.Context, id string) (*Tournament, error)
}
