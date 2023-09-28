package commands

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/domain"
)

type RemoveStore struct {
	ID string
}

type RemoveStoreHandler struct {
	es domain.StoreEventStore
}

func NewRemoveStoreHandler(es domain.StoreEventStore) RemoveStoreHandler {
	return RemoveStoreHandler{es: es}
}

func (h RemoveStoreHandler) RemoveStore(ctx context.Context, cmd RemoveStore) error {

	store, err := h.es.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = store.Remove()
	if err != nil {
		return err
	}

	return h.es.Save(ctx, store)
}
