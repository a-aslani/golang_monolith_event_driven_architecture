package monolith

import (
	"context"
	"database/sql"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/config"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/waiter"
	"github.com/go-chi/chi/v5"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Monolith interface {
	Config() config.AppConfig
	DB() *sql.DB
	ESDB() *esdb.Client
	JS() nats.JetStreamContext
	Logger() zerolog.Logger
	Mux() *chi.Mux
	RPC() *grpc.Server
	Waiter() waiter.Waiter
}

type Module interface {
	Startup(context.Context, Monolith) error
}
