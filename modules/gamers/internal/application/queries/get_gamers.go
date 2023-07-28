package queries

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain"
)

type GetGamers struct{}

type GetGamersHandler struct {
	repo domain.GamerRepository
}

func NewGetGamersHandler(repo domain.GamerRepository) GetGamersHandler {
	return GetGamersHandler{repo: repo}
}

func (h GetGamersHandler) GetGamers(ctx context.Context, query GetGamers) ([]*domain.GamerModel, error) {
	return h.repo.FindGamers(ctx)
}
