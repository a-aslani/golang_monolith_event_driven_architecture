package commands

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/domain"
)

type EditStore struct {
	ID     string
	Name   string
	Amount int
	Price  float64
}

type EditStoreHandler struct {
	eventStore domain.StoreEventStore
}

func NewEditStoreHandler(eventStore domain.StoreEventStore) EditStoreHandler {
	return EditStoreHandler{eventStore: eventStore}
}

func (h EditStoreHandler) EditStore(ctx context.Context, cmd EditStore) error {

	store, err := h.eventStore.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = store.Edit(cmd.Name, cmd.Amount, cmd.Price)
	if err != nil {
		return err
	}

	return h.eventStore.Save(ctx, store)
}
