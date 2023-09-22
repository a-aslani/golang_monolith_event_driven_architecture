package domain

import "github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain/value_objects"

const (
	GamerCreatedEvent     = "V1.Gamers.GamerCreated"
	GamerApprovedEvent    = "V1.Gamers.GamerApproved"
	GamerDisapprovedEvent = "V1.Gamers.GamerDisapproved"
	GamerUpdatedGemEvent  = "V1.Gamers.UpdatedGem"
)

type GamerCreated struct {
	FullName   value_objects.GamerName
	Email      value_objects.GamerEmail
	Password   value_objects.GamerPassword
	IsApproved bool
}

func (GamerCreated) Key() string { return GamerCreatedEvent }

type GamerApproved struct{}

func (GamerApproved) Key() string { return GamerApprovedEvent }

type GamerDisapproved struct{}

func (GamerDisapproved) Key() string { return GamerDisapprovedEvent }

type GamerUpdatedGem struct {
	Amount value_objects.GamerGem
}

func (GamerUpdatedGem) Key() string {
	return GamerUpdatedGemEvent
}
