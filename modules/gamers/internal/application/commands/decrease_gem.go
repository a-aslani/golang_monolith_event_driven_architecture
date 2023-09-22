package commands

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain"
)

type DecreaseGem struct {
	ID     string
	Amount int
}

type DecreaseGemHandler struct {
	eventStore domain.GamerEventStore
}

func NewDecreaseGemHandler(eventStore domain.GamerEventStore) DecreaseGemHandler {
	return DecreaseGemHandler{
		eventStore: eventStore,
	}
}

func (h DecreaseGemHandler) DecreaseGem(ctx context.Context, cmd DecreaseGem) error {

	gamer, err := h.eventStore.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = gamer.DecreaseGem(cmd.Amount)
	if err != nil {
		return err
	}

	return h.eventStore.Save(ctx, gamer)
}
