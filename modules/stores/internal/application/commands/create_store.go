package commands

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/domain"
)

type CreateStore struct {
	ID     string
	Name   string
	Amount int
	Price  float64
}

type CreateStoreHandler struct {
	eventStore domain.StoreEventStore
}

func NewCreateStoreHandler(eventStore domain.StoreEventStore) CreateStoreHandler {
	return CreateStoreHandler{eventStore: eventStore}
}

func (h CreateStoreHandler) CreateStore(ctx context.Context, cmd CreateStore) error {

	store, err := domain.CreateStore(cmd.ID, cmd.Name, cmd.Amount, cmd.Price)
	if err != nil {
		return err
	}

	return h.eventStore.Save(ctx, store)
}
