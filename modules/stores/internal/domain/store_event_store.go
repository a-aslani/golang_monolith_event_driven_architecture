package domain

import "context"

type StoreEventStore interface {
	Load(ctx context.Context, id string) (*Store, error)
	Save(ctx context.Context, store *Store) error
}
