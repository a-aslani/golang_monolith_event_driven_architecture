package gamers

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/application"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/endpoint/grpc"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/endpoint/rest"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/handlers"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/infrastructure/logging"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/infrastructure/postgres"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/infrastructure/utils"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/es"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/es/aggregate_store"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/monolith"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/registry"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/registry/serdes"
)

type Module struct {
}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {

	// setup Driven adapters
	reg := registry.New()
	err := registrations(reg)
	if err != nil {
		return err
	}
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	aggregateGamer := es.AggregateStoreWithMiddleware(
		aggregate_store.NewEventStoreDB(mono.ESDB(), reg),
		es.NewEventPublisher(domainDispatcher),
		//pg.NewSnapshotStore("gamers.snapshots", mono.DB(), reg),
	)
	gamerEventStore := es.NewAggregateRepository[*domain.Gamer](domain.GamerAggregate, reg, aggregateGamer)
	gamerRepo := postgres.NewGamerRepository(mono.DB())
	u := utils.NewUtils()

	// setup application
	app := logging.LogApplicationAccess(
		application.New(gamerEventStore, gamerRepo, u),
		mono.Logger(),
	)
	gamerHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewGamerHandlers(gamerRepo),
		"Gamer", mono.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	handlers.RegisterGamerHandlers(gamerHandlers, domainDispatcher)

	return nil
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Gamer
	if err = serde.Register(domain.Gamer{}, func(v any) error {
		store := v.(*domain.Gamer)
		store.Aggregate = es.NewAggregate("", domain.GamerAggregate)
		return nil
	}); err != nil {
		return
	}

	// gamer events
	if err = serde.Register(domain.GamerCreated{}); err != nil {
		return
	}
	if err = serde.Register(domain.GamerApproved{}); err != nil {
		return
	}
	if err = serde.Register(domain.GamerDisapproved{}); err != nil {
		return
	}

	// gamer snapshots
	if err = serde.RegisterKey(domain.GamerV1{}.SnapshotName(), domain.GamerV1{}); err != nil {
		return
	}

	return
}
