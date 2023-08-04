package handlers

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/gamerspb"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/am"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
)

func RegisterGamerHandlers(gamerHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {

	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, msg am.EventMessage) error {
		return gamerHandlers.HandleEvent(ctx, msg)
	})

	return stream.Subscribe(gamerspb.GamerAggregateChannel, evtMsgHandler, am.MessageFilter{
		gamerspb.GamerCreatedEvent,
	})
}
