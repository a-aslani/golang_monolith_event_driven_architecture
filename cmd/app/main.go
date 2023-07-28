package main

import (
	"database/sql"
	"fmt"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/config"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/event_strore_db"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/logger"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/monolith"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/rpc"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/waiter"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/web"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
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
	m.esdb, err = event_strore_db.NewEventStoreDB(cfg.ESDB.Conn)
	if err != nil {
		return err
	}
	defer func(esdb *esdb.Client) {
		err := esdb.Close()
		if err != nil {
			return
		}
	}(m.esdb)

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
	m.logger = logger.New(logger.LogConfig{
		Environment: cfg.Environment,
		LogLevel:    logger.Level(cfg.LogLevel),
	})
	m.rpc = initRpc(cfg.Rpc)
	m.mux = initMux(cfg.Web)
	m.waiter = waiter.New(waiter.CatchSignals())

	// init modules
	m.modules = []monolith.Module{
		&gamers.Module{},
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

func initRpc(_ rpc.RpcConfig) *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)

	return server
}

func initMux(_ web.WebConfig) *chi.Mux {
	return chi.NewMux()
}
