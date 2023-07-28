package grpc

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/gamerspb"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/application"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/application/commands"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	gamerspb.UnimplementedGamersServiceServer
}

var _ gamerspb.GamersServiceServer = (*server)(nil)

func RegisterServer(ctx context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	gamerspb.RegisterGamersServiceServer(registrar, server{app: app})
	return nil
}

func (s server) DisapproveGamer(ctx context.Context, request *gamerspb.DisapproveGamerRequest) (*gamerspb.DisapproveGamerResponse, error) {

	err := s.app.DisapproveGamer(
		ctx,
		commands.DisapproveGamer{
			ID: request.GetId(),
		},
	)
	if err != nil {
		return nil, err
	}

	return &gamerspb.DisapproveGamerResponse{
		Id: request.GetId(),
	}, nil

}

func (s server) CreateGamer(ctx context.Context, request *gamerspb.CreateGamerRequest) (*gamerspb.CreateGamerResponse, error) {

	gamerID := uuid.New().String()

	err := s.app.CreateGamer(ctx, commands.CreateGamer{
		ID:        gamerID,
		FirstName: request.GetFirstName(),
		LastName:  request.GetLastName(),
		Email:     request.GetEmail(),
		Password:  request.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &gamerspb.CreateGamerResponse{
		Id: gamerID,
	}, nil
}
