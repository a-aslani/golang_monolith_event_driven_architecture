package logging

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/application"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/application/commands"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/application/queries"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/domain"
	"github.com/rs/zerolog"
)

type Application struct {
	app    application.App
	logger zerolog.Logger
}

var _ application.App = (*Application)(nil)

func LogApplicationAccess(app application.App, logger zerolog.Logger) Application {
	return Application{
		app:    app,
		logger: logger,
	}
}

func (a Application) CreateStore(ctx context.Context, cmd commands.CreateStore) (err error) {
	a.logger.Info().Msg("--> Stores.CreateStore")
	defer func() { a.logger.Info().Err(err).Msg("<-- Stores.CreateStore") }()
	return a.app.CreateStore(ctx, cmd)
}

func (a Application) EditStore(ctx context.Context, cmd commands.EditStore) (err error) {
	a.logger.Info().Msg("--> Stores.EditStore")
	defer func() { a.logger.Info().Err(err).Msg("<-- Stores.EditStore") }()
	return a.app.EditStore(ctx, cmd)
}

func (a Application) RemoveStore(ctx context.Context, cmd commands.RemoveStore) (err error) {
	a.logger.Info().Msg("--> Stores.RemoveStore")
	defer func() { a.logger.Info().Err(err).Msg("<-- Stores.RemoveStore") }()
	return a.app.RemoveStore(ctx, cmd)
}

func (a Application) GetStore(ctx context.Context, query queries.GetStore) (_ *domain.StoreDTO, err error) {
	a.logger.Info().Msg("--> Stores.GetStore")
	defer func() { a.logger.Info().Err(err).Msg("<-- Stores.GetStore") }()
	return a.app.GetStore(ctx, query)
}

func (a Application) GetStores(ctx context.Context, query queries.GetStores) (_ []*domain.StoreDTO, err error) {
	a.logger.Info().Msg("--> Stores.GetStores")
	defer func() { a.logger.Info().Err(err).Msg("<-- Stores.GetStores") }()
	return a.app.GetStores(ctx, query)
}
