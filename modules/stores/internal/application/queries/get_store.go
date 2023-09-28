package queries

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/domain"
)

type GetStore struct {
	ID string
}

type GetStoreHandler struct {
	repo domain.StoreRepository
}

func NewGetStore(repo domain.StoreRepository) GetStoreHandler {
	return GetStoreHandler{repo: repo}
}

func (h GetStoresHandler) GetStore(ctx context.Context, query GetStore) (*domain.StoreDTO, error) {
	return h.repo.Find(ctx, query.ID)
}
