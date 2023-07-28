package logging

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/application"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/application/commands"
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
