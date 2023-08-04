package notifications

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/gamerspb"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/notifications/internal/application"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/notifications/internal/handlers"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/notifications/internal/logging"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/am"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/jetstream"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/monolith"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/registry"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {

	// setup Driven adapters
	reg := registry.New()
	if err := gamerspb.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS()))

	// setup application

	gamerHandlers := logging.LogEventHandlerAccess[ddd.Event](
		application.NewGamerHandlers(mono.Logger()),
		"Gamer", mono.Logger(),
	)

	// setup Driver adapters
	if err := handlers.RegisterGamerHandlers(gamerHandlers, eventStream); err != nil {
		return err
	}

	return nil
}
