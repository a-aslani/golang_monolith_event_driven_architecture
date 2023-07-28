package queries

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain"
)

type GetGamer struct {
	ID string
}

type GetGamerHandler struct {
	repo domain.GamerRepository
}

func NewGetGamerHandler(repo domain.GamerRepository) GetGamerHandler {
	return GetGamerHandler{
		repo: repo,
	}
}

func (h GetGamerHandler) GetGamer(ctx context.Context, query GetGamer) (*domain.GamerModel, error) {
	return h.repo.FindGamer(ctx, query.ID)
}
