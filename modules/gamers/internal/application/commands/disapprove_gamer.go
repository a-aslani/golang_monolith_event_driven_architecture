package commands

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain"
)

type DisapproveGamer struct {
	ID string
}

type DisapproveGamerHandler struct {
	eventStore domain.GamerEventStore
}

func NewDisapproveGamerHandler(eventStore domain.GamerEventStore) DisapproveGamerHandler {
	return DisapproveGamerHandler{
		eventStore: eventStore,
	}
}

func (h DisapproveGamerHandler) DisapproveGamer(ctx context.Context, cmd DisapproveGamer) error {

	gamer, err := h.eventStore.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	gamer.Disapprove()

	return h.eventStore.Save(ctx, gamer)
}
