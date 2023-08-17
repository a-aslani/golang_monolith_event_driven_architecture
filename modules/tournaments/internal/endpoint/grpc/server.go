package grpc

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/application"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/internal/application/commands"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/tournaments/tournamentspb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"time"
)

type server struct {
	app application.App
	tournamentspb.UnimplementedTournamentsServiceServer
}

var _ tournamentspb.TournamentsServiceServer = (*server)(nil)

func RegisterServer(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {

	tournamentspb.RegisterTournamentsServiceServer(registrar, server{
		app: app,
	})

	return nil
}

func (s server) CreateTournament(ctx context.Context, req *tournamentspb.CreateTournamentRequest) (*tournamentspb.CreateTournamentResponse, error) {

	tournamentID := uuid.New().String()

	err := s.app.CreateTournament(ctx, commands.CreateTournament{
		ID:          tournamentID,
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Gamer1ID:    req.GetGamer_1Id(),
		Gamer2ID:    req.GetGamer_2Id(),
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &tournamentspb.CreateTournamentResponse{Id: tournamentID}, err
}
