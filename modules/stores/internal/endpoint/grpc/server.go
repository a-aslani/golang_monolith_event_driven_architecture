package grpc

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/application"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/application/commands"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/application/queries"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/domain"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/storespb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	storespb.UnimplementedStoresServiceServer
}

var _ storespb.StoresServiceServer = (*server)(nil)

func RegisterServer(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	storespb.RegisterStoresServiceServer(registrar, server{app: app})
	return nil
}

func (s server) CreateStore(ctx context.Context, req *storespb.CreateStoreRequest) (*storespb.CreateStoreResponse, error) {

	id := uuid.New().String()

	err := s.app.CreateStore(ctx, commands.CreateStore{
		ID:     id,
		Name:   req.GetName(),
		Amount: int(req.GetAmount()),
		Price:  float64(req.GetPrice()),
	})
	if err != nil {
		return nil, err
	}

	return &storespb.CreateStoreResponse{Id: id}, nil
}

func (s server) EditStore(ctx context.Context, req *storespb.EditStoreRequest) (*storespb.EditStoreResponse, error) {

	err := s.app.EditStore(ctx, commands.EditStore{
		ID:     req.GetId(),
		Name:   req.GetName(),
		Amount: int(req.GetAmount()),
		Price:  float64(req.GetPrice()),
	})
	if err != nil {
		return nil, err
	}

	return &storespb.EditStoreResponse{Id: req.GetId()}, nil
}

func (s server) RemoveStore(ctx context.Context, req *storespb.RemoveStoreRequest) (*storespb.RemoveStoreResponse, error) {

	err := s.app.RemoveStore(ctx, commands.RemoveStore{ID: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &storespb.RemoveStoreResponse{}, err
}

func (s server) GetStore(ctx context.Context, req *storespb.GetStoreRequest) (*storespb.GetStoreResponse, error) {

	store, err := s.app.GetStore(ctx, queries.GetStore{ID: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &storespb.GetStoreResponse{Store: s.storeFromDomain(store)}, err
}

func (s server) GetStores(ctx context.Context, req *storespb.GetStoresRequest) (*storespb.GetStoresResponse, error) {

	stores, err := s.app.GetStores(ctx, queries.GetStores{})
	if err != nil {
		return nil, err
	}

	data := make([]*storespb.Store, 0)

	for _, item := range stores {
		data = append(data, s.storeFromDomain(item))
	}

	return &storespb.GetStoresResponse{Stores: data}, nil
}

func (s server) storeFromDomain(store *domain.StoreDTO) *storespb.Store {
	return &storespb.Store{
		Id:     store.ID,
		Name:   store.Name,
		Amount: int32(store.Amount),
		Price:  float32(store.Price),
	}
}
