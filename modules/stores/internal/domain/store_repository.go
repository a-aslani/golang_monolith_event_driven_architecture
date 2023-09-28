package domain

import "context"

type StoreRepository interface {
	FindAll(ctx context.Context) ([]*StoreDTO, error)
	Find(ctx context.Context, id string) (*StoreDTO, error)
	Insert(ctx context.Context, id string, name string, amount int, price float64) error
	Update(ctx context.Context, id string, name string, amount int, price float64) error
	Remove(ctx context.Context, id string) error
}

type StoreDTO struct {
	ID     string
	Name   string
	Amount int
	Price  float64
}
