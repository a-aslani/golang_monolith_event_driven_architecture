package main

import (
	"database/sql"
	"fmt"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/config"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/notifications"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/logger"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/monolith"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/rpc"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/waiter"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/web"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	var cfg config.AppConfig
	// parse config/env/...
	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}

	m := app{cfg: cfg}

	// init infrastructure...

	// init EventStore db
	m.esdb, err = initEventStoreDB(cfg.ESDB.Conn)
	if err != nil {
		return err
	}
	defer func(esdb *esdb.Client) {
		err := esdb.Close()
		if err != nil {
			return
		}
	}(m.esdb)

	// connect to the Postgres
	m.db, err = sql.Open("pgx", cfg.PG.Conn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(m.db)

	// init nats & JetStream
	m.nc, err = nats.Connect(cfg.Nats.URL)
	if err != nil {
		return err
	}
	defer m.nc.Close()
	m.js, err = initJetStream(cfg.Nats, m.nc)
	if err != nil {
		return err
	}

	m.logger = initLogger(cfg)
	m.rpc = initRpc(cfg.Rpc)
	m.mux = initMux(cfg.Web)
	m.waiter = waiter.New(waiter.CatchSignals())

	// init modules
	m.modules = []monolith.Module{
		&gamers.Module{},
		&notifications.Module{},
	}

	if err = m.startupModules(); err != nil {
		return err
	}

	// Mount general web resources
	m.mux.Mount("/", http.FileServer(http.FS(web.WebUI)))

	fmt.Println("Started Event-Driven application")
	defer fmt.Println("Stopped Event-Driven application")

	m.waiter.Add(
		m.waitForWeb,
		m.waitForRPC,
	)

	return m.waiter.Wait()
}

func initLogger(cfg config.AppConfig) zerolog.Logger {
	return logger.New(logger.LogConfig{
		Environment: cfg.Environment,
		LogLevel:    logger.Level(cfg.LogLevel),
	})
}

func initRpc(_ rpc.RpcConfig) *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)

	return server
}

func initMux(_ web.WebConfig) *chi.Mux {
	return chi.NewMux()
}

func initJetStream(cfg config.NatsConfig, nc *nats.Conn) (nats.JetStreamContext, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     cfg.Stream,
		Subjects: []string{fmt.Sprintf("%s.>", cfg.Stream)},
	})

	return js, err
}

func initEventStoreDB(connectionString string) (*esdb.Client, error) {
	settings, err := esdb.ParseConnectionString(connectionString)
	if err != nil {
		return nil, err
	}
	return esdb.NewClient(settings)
}
