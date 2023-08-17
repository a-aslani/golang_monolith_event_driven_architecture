package logging

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/application"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/application/commands"
	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

var _ application.App = (*Application)(nil)

func LogApplicationAccess(app application.App, logger zerolog.Logger) Application {
	return Application{
		App:    app,
		logger: logger,
	}
}

func (a Application) CreateTournament(ctx context.Context, cmd commands.CreateTournament) error {
	a.logger.Info().Msg("--> tournaments.CreateTournament")
	defer func() {
		a.logger.Info().Msg("<-- tournaments.CreateTournament")
	}()
	return a.App.CreateTournament(ctx, cmd)
}
