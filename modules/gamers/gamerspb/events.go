package gamerspb

import (
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/registry"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/registry/serdes"
)

const (
	GamerAggregateChannel = "app.gamers.events.Gamer"

	GamerCreatedEvent = "gamersapi.GamerCreated"
)

func (*GamerCreated) Key() string { return GamerCreatedEvent }

func Registrations(reg registry.Registry) error {
	serde := serdes.NewProtoSerde(reg)

	// Gamer events
	if err := serde.Register(&GamerCreated{}); err != nil {
		return err
	}

	return nil
}
