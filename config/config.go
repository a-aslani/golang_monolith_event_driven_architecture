package config

import (
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/rpc"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/web"
	"github.com/kelseyhightower/envconfig"
	"github.com/stackus/dotenv"
	"os"
	"time"
)

type (
	PGConfig struct {
		Conn string `required:"true"`
	}

	ESDBConfig struct {
		Conn string `required:"true"`
	}

	AppConfig struct {
		Environment     string
		LogLevel        string `envconfig:"LOG_LEVEL" default:"DEBUG"`
		PG              PGConfig
		ESDB            ESDBConfig
		Rpc             rpc.RpcConfig
		Web             web.WebConfig
		ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
	}
)

func InitConfig() (cfg AppConfig, err error) {
	if err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
		return
	}
	err = envconfig.Process("", &cfg)
	return
}
