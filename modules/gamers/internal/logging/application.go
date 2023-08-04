package logging

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/application"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/application/commands"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/application/queries"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain"
	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

var _ application.App = (*Application)(nil)

func LogApplicationAccess(application application.App, logger zerolog.Logger) Application {
	return Application{
		App:    application,
		logger: logger,
	}
}

func (a Application) CreateGamer(ctx context.Context, cmd commands.CreateGamer) (err error) {
	a.logger.Info().Msg("--> Gamers.CreateGamer")
	defer func() { a.logger.Info().Err(err).Msg("<-- Gamers.CreateGamers") }()
	return a.App.CreateGamer(ctx, cmd)
}

func (a Application) DisapproveGamer(ctx context.Context, cmd commands.DisapproveGamer) (err error) {
	a.logger.Info().Msg("--> Gamers.DisapproveGamer")
	defer func() { a.logger.Info().Err(err).Msg("<-- Gamers.DisapproveGamer") }()
	return a.App.DisapproveGamer(ctx, cmd)
}

func (a Application) GetGamer(ctx context.Context, query queries.GetGamer) (_ *domain.GamerDTO, err error) {
	a.logger.Info().Msg("--> Gamers.GetGamer")
	defer func() { a.logger.Info().Err(err).Msg("<-- Gamers.GetGamer") }()
	return a.App.GetGamer(ctx, query)
}

func (a Application) GetGamers(ctx context.Context, query queries.GetGamers) (_ []*domain.GamerDTO, err error) {
	a.logger.Info().Msg("--> Gamers.GetGamers")
	defer func() { a.logger.Info().Err(err).Msg("<-- Gamers.GetGamers") }()
	return a.App.GetGamers(ctx, query)
}
