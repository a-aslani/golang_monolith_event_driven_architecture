package commands

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain"
)

type CreateGamer struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type CreateGamerHandler struct {
	eventStore domain.GamerEventStore
	utils      domain.Utils
}

func NewCreateGamerHandler(eventStore domain.GamerEventStore, utils domain.Utils) CreateGamerHandler {
	return CreateGamerHandler{
		eventStore: eventStore,
		utils:      utils,
	}
}

func (h CreateGamerHandler) CreateGamer(ctx context.Context, cmd CreateGamer) error {

	hashedPassword, err := h.utils.Hash(cmd.Password)
	if err != nil {
		return err
	}

	gamer, err := domain.CreateGamer(cmd.ID, cmd.FirstName, cmd.LastName, cmd.Email, hashedPassword)
	if err != nil {
		return err
	}

	return h.eventStore.Save(ctx, gamer)
}
