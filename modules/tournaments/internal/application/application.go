package application

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/application/commands"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/domain"
)

type App interface {
	Commands
	Queries
}

type Commands interface {
	CreateTournament(ctx context.Context, cmd commands.CreateTournament) error
}

type Queries interface {
}

type Application struct {
	appCommands
	appQueries
}

type appCommands struct {
	commands.CreateTournamentHandler
}

type appQueries struct {
}

var _ App = (*Application)(nil)

func New(tournamentEventStore domain.TournamentEventStore) *Application {
	return &Application{
		appCommands: appCommands{
			CreateTournamentHandler: commands.NewCreateTournamentHandler(tournamentEventStore),
		},
		appQueries: appQueries{},
	}
}
