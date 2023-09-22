package grpc

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/gamerspb"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/application"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/application/commands"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/application/queries"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	gamerspb.UnimplementedGamersServiceServer
}

var _ gamerspb.GamersServiceServer = (*server)(nil)

func RegisterServer(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	gamerspb.RegisterGamersServiceServer(registrar, server{app: app})
	return nil
}

func (s server) DecreaseGem(ctx context.Context, req *gamerspb.DecreaseGemRequest) (*gamerspb.DecreaseGemResponse, error) {

	err := s.app.DecreaseGem(ctx, commands.DecreaseGem{
		ID:     req.GetId(),
		Amount: int(req.GetAmount()),
	})
	if err != nil {
		return nil, err
	}

	return &gamerspb.DecreaseGemResponse{Id: req.GetId()}, nil

}

func (s server) IncreaseGem(ctx context.Context, req *gamerspb.IncreaseGemRequest) (*gamerspb.IncreaseGemResponse, error) {

	err := s.app.IncreaseGem(ctx, commands.IncreaseGem{
		ID:     req.GetId(),
		Amount: int(req.GetAmount()),
	})
	if err != nil {
		return nil, err
	}

	return &gamerspb.IncreaseGemResponse{Id: req.GetId()}, nil
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

func (s server) GetGamer(ctx context.Context, request *gamerspb.GetGamerRequest) (*gamerspb.GetGamerResponse, error) {

	gamer, err := s.app.GetGamer(ctx, queries.GetGamer{ID: request.GetId()})
	if err != nil {
		return nil, err
	}

	return &gamerspb.GetGamerResponse{Gamer: s.gamerFromDomain(gamer)}, err
}

func (s server) GetGamers(ctx context.Context, request *gamerspb.GetGamersRequest) (*gamerspb.GetGamersResponse, error) {

	gamers, err := s.app.GetGamers(ctx, queries.GetGamers{})
	if err != nil {
		return nil, err
	}

	protoGamers := make([]*gamerspb.Gamer, 0)

	for _, item := range gamers {
		protoGamers = append(protoGamers, s.gamerFromDomain(item))
	}

	return &gamerspb.GetGamersResponse{Gamers: protoGamers}, err
}

func (s server) gamerFromDomain(gamer *domain.GamerDTO) *gamerspb.Gamer {
	return &gamerspb.Gamer{
		Id:         gamer.ID,
		FirstName:  gamer.FirstName,
		LastName:   gamer.LastName,
		Email:      gamer.Email,
		IsApproved: gamer.IsApproved,
	}
}
