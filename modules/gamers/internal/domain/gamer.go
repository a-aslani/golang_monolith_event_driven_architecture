package domain

import (
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain/value_objects"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/es"
	"github.com/stackus/errors"
)

const GamerAggregate = "gamers.Gamer"

type Gamer struct {
	es.Aggregate
	name       value_objects.GamerName
	email      value_objects.GamerEmail
	password   value_objects.GamerPassword
	isApproved bool
	gem        value_objects.GamerGem
}

var _ interface {
	es.EventApplier
	es.Snapshotter
} = (*Gamer)(nil)

func NewGamer(id string) *Gamer {

	return &Gamer{
		Aggregate: es.NewAggregate(id, GamerAggregate),
	}
}

func CreateGamer(id, fistName, lastName, email, password string) (*Gamer, error) {

	gamerName, err := value_objects.NewGamerName(fistName, lastName)
	if err != nil {
		return nil, err
	}

	gamerEmail, err := value_objects.NewGamerEmail(email)
	if err != nil {
		return nil, err
	}

	gamerPassword, err := value_objects.NewGamerPassword(password)
	if err != nil {
		return nil, err
	}

	gamer := NewGamer(id)

	gamer.AddEvent(GamerCreatedEvent, &GamerCreated{
		FullName:   gamerName,
		Email:      gamerEmail,
		Password:   gamerPassword,
		IsApproved: false,
	})

	gamer.AddEvent(GamerApprovedEvent, &GamerApproved{})

	return gamer, nil
}

func (a *Gamer) Disapprove() {
	a.AddEvent(GamerDisapprovedEvent, &GamerDisapproved{})
}

func (a *Gamer) IncreaseGem(amount int) error {

	gem, err := value_objects.NewGamerGem(a.gem.Value + amount)
	if err != nil {
		return err
	}

	a.AddEvent(GamerUpdatedGemEvent, &GamerUpdatedGem{Amount: gem})

	return nil
}

func (a *Gamer) DecreaseGem(amount int) error {

	gem, err := value_objects.NewGamerGem(a.gem.Value - amount)
	if err != nil {
		return err
	}

	a.AddEvent(GamerUpdatedGemEvent, &GamerUpdatedGem{Amount: gem})

	return err
}

func (Gamer) Key() string { return GamerAggregate }

func (a *Gamer) ApplyEvent(event ddd.Event) error {

	switch payload := event.Payload().(type) {
	case *GamerCreated:
		a.name = payload.FullName
		a.email = payload.Email
		a.password = payload.Password
		a.isApproved = payload.IsApproved

	case *GamerApproved:
		a.isApproved = true

	case *GamerDisapproved:
		a.isApproved = false

	case *GamerUpdatedGem:
		a.gem = payload.Amount

	default:
		return errors.ErrInternal.Msgf("%T received the event %s with unexpected payload %T", a, event.EventName(), payload)
	}

	return nil
}

func (a *Gamer) ApplySnapshot(snapshot es.Snapshot) error {

	switch ss := snapshot.(type) {
	case *GamerV1:
		a.name = ss.Name
		a.email = ss.Email
		a.password = ss.Password
		a.isApproved = ss.IsApproved
		a.gem = ss.Gem
	default:
		return errors.ErrInternal.Msgf("%T received the unexpected snapshot %T", a, snapshot)
	}

	return nil
}

func (a *Gamer) ToSnapshot() es.Snapshot {
	return GamerV1{
		Name:       a.name,
		Email:      a.email,
		Password:   a.password,
		IsApproved: a.isApproved,
		Gem:        a.gem,
	}
}
