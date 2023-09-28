package queries

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/domain"
)

type GetStores struct{}

type GetStoresHandler struct {
	repo domain.StoreRepository
}

func NewGetStores(repo domain.StoreRepository) GetStoresHandler {
	return GetStoresHandler{repo: repo}
}

func (h GetStoresHandler) GetStores(ctx context.Context, query GetStores) ([]*domain.StoreDTO, error) {
	return h.repo.FindAll(ctx)
}
