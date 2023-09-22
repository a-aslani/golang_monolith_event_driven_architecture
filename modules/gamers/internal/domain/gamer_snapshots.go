package domain

import (
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain/value_objects"
)

type GamerV1 struct {
	Name       value_objects.GamerName
	Email      value_objects.GamerEmail
	Password   value_objects.GamerPassword
	IsApproved bool
	Gem        value_objects.GamerGem
}

func (g GamerV1) SnapshotName() string {
	return "gamers.GamerV1"
}
