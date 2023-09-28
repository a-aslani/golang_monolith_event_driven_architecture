package application

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/application/commands"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/application/queries"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/domain"
)

type Commands interface {
	CreateStore(ctx context.Context, cmd commands.CreateStore) error
	EditStore(ctx context.Context, cmd commands.EditStore) error
	RemoveStore(ctx context.Context, cmd commands.RemoveStore) error
}

type Queries interface {
	GetStore(ctx context.Context, query queries.GetStore) (*domain.StoreDTO, error)
	GetStores(ctx context.Context, query queries.GetStores) ([]*domain.StoreDTO, error)
}

type App interface {
	Commands
	Queries
}

type appCommands struct {
	commands.RemoveStoreHandler
	commands.EditStoreHandler
	commands.CreateStoreHandler
}

type appQueries struct {
	queries.GetStoresHandler
	queries.GetStoreHandler
}

type Application struct {
	appCommands
	appQueries
}

var _ App = (*Application)(nil)

func New(es domain.StoreEventStore, repo domain.StoreRepository) *Application {
	return &Application{
		appCommands: appCommands{
			RemoveStoreHandler: commands.NewRemoveStoreHandler(es),
			EditStoreHandler:   commands.NewEditStoreHandler(es),
			CreateStoreHandler: commands.NewCreateStoreHandler(es),
		},
		appQueries: appQueries{
			GetStoresHandler: queries.NewGetStores(repo),
			GetStoreHandler:  queries.NewGetStore(repo),
		},
	}
}
