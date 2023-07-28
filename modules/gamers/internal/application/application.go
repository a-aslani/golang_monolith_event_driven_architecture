package application

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/application/commands"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/application/queries"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain"
)

type App interface {
	Commands
	Queries
}

type Commands interface {
	CreateGamer(ctx context.Context, cmd commands.CreateGamer) error
	DisapproveGamer(ctx context.Context, cmd commands.DisapproveGamer) error
}

type Queries interface {
	GetGamer(ctx context.Context, query queries.GetGamer) (*domain.GamerDTO, error)
	GetGamers(ctx context.Context, query queries.GetGamers) ([]*domain.GamerDTO, error)
}

type Application struct {
	appCommands
	appQueries
}

type appCommands struct {
	commands.CreateGamerHandler
	commands.DisapproveGamerHandler
}

type appQueries struct {
	queries.GetGamersHandler
	queries.GetGamerHandler
}

var _ App = (*Application)(nil)

func New(gamerEventStore domain.GamerEventStore, gamerRepo domain.GamerRepository, utils domain.Utils) *Application {
	return &Application{
		appCommands: appCommands{
			CreateGamerHandler:     commands.NewCreateGamerHandler(gamerEventStore, utils),
			DisapproveGamerHandler: commands.NewDisapproveGamerHandler(gamerEventStore),
		},
		appQueries: appQueries{
			GetGamerHandler:  queries.NewGetGamerHandler(gamerRepo),
			GetGamersHandler: queries.NewGetGamersHandler(gamerRepo),
		},
	}
}
