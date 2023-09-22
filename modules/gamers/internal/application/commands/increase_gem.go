package commands

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain"
)

type IncreaseGem struct {
	ID     string
	Amount int
}

type IncreaseGemHandler struct {
	eventStore domain.GamerEventStore
}

func NewIncreaseGemHandler(eventStore domain.GamerEventStore) IncreaseGemHandler {
	return IncreaseGemHandler{eventStore: eventStore}
}

func (h IncreaseGemHandler) IncreaseGem(ctx context.Context, cmd IncreaseGem) error {

	gamer, err := h.eventStore.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = gamer.IncreaseGem(cmd.Amount)
	if err != nil {
		return err
	}

	return h.eventStore.Save(ctx, gamer)
}
