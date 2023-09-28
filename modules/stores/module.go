package stores

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/application"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/domain"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/endpoint/grpc"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/endpoint/rest"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/handlers"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/infrastructure/postgres"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/logging"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/es"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/es/aggregate_store"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/monolith"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/registry"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/registry/serdes"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {

	// setup Driven adapters
	reg := registry.New()
	err := registrations(reg)
	if err != nil {
		return err
	}

	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	aggregateStore := es.AggregateStoreWithMiddleware(
		aggregate_store.NewEventStoreDB(mono.ESDB(), reg, mono.Logger()),
		es.NewEventPublisher(domainDispatcher),
		//pg.NewSnapshotStore("stores.snapshots", mono.DB(), reg),
	)
	storeEventStore := es.NewAggregateRepository[*domain.Store](domain.StoreAggregate, reg, aggregateStore)
	storeRepo := postgres.NewStoreRepository(mono.DB())

	// setup application
	app := logging.LogApplicationAccess(
		application.New(storeEventStore, storeRepo),
		mono.Logger(),
	)
	storeHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewStoreHandlers(storeRepo),
		"Store", mono.Logger(),
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
	handlers.RegisterStoreHandlers(storeHandlers, domainDispatcher)

	return nil
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Store
	if err = serde.Register(domain.Store{}, func(v any) error {
		store := v.(*domain.Store)
		store.Aggregate = es.NewAggregate("", domain.StoreAggregate)
		return nil
	}); err != nil {
		return
	}

	// Store events
	if err = serde.Register(domain.StoreCreated{}); err != nil {
		return
	}
	if err = serde.Register(domain.StoreEdited{}); err != nil {
		return
	}
	if err = serde.Register(domain.StoreRemoved{}); err != nil {
		return
	}

	// Store snapshots
	if err = serde.RegisterKey(domain.StoreV1{}.SnapshotName(), domain.StoreV1{}); err != nil {
		return
	}

	return
}
