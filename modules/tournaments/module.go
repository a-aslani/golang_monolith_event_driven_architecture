package tournaments

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/application"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/domain"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/endpoint/grpc"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/endpoint/rest"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/handlers"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/infrastructure/postgres"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/logging"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/es"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/es/aggregate_store"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/monolith"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/registry"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/registry/serdes"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {

	// setup driven adapters
	reg := registry.New()
	err := registrations(reg)
	if err != nil {
		return err
	}

	//eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS()))
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	aggregateTournamentStore := es.AggregateStoreWithMiddleware(
		aggregate_store.NewEventStoreDB(mono.ESDB(), reg, mono.Logger()),
		es.NewEventPublisher(domainDispatcher),
	)

	tournamentEventStore := es.NewAggregateRepository[*domain.Tournament](domain.TournamentAggregate, reg, aggregateTournamentStore)
	tournamentRepo := postgres.NewTournamentRepository(mono.DB())

	// setup application
	app := logging.LogApplicationAccess(
		application.New(tournamentEventStore),
		mono.Logger(),
	)
	tournamentHandlers := handlers.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewTournamentHandlers(tournamentRepo),
		"Tournament",
		mono.Logger(),
	)

	// setup driven adapters
	if err = grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err = rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err = rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	handlers.RegisterTournamentHandlers(tournamentHandlers, domainDispatcher)

	return nil
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// tournament
	if err = serde.Register(domain.Tournament{}, func(v interface{}) error {

		tournament := v.(*domain.Tournament)
		tournament.Aggregate = es.NewAggregate("", domain.TournamentAggregate)
		return nil

	}); err != nil {
		return err
	}

	// tournament events
	if err = serde.Register(domain.TournamentCreated{}); err != nil {
		return err
	}

	// tournament snapshot
	if err = serde.RegisterKey(domain.TournamentV1{}.SnapshotName(), domain.TournamentV1{}); err != nil {
		return err
	}

	return
}
