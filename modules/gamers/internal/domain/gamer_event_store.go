package domain

import "context"

type GamerEventStore interface {
	Load(ctx context.Context, id string) (*Gamer, error)
	Save(ctx context.Context, gamer *Gamer) error
}
