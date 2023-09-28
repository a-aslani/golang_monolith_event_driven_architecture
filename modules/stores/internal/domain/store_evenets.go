package domain

import "github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/domain/value_objects"

const (
	StoreCreatedEvent = "V1.Stores.StoreCreated"
	StoreEditedEvent  = "V1.Stores.StoreEdited"
	StoreRemovedEvent = "V1.Stores.StoreRemoved"
)

type StoreRemoved struct{}

func (s StoreRemoved) Key() string {
	return StoreRemovedEvent
}

type StoreEdited struct {
	Name   value_objects.StoreName
	Amount value_objects.StoreAmount
	Price  value_objects.StorePrice
}

func (s StoreEdited) Key() string {
	return StoreEditedEvent
}

type StoreCreated struct {
	Name   value_objects.StoreName
	Amount value_objects.StoreAmount
	Price  value_objects.StorePrice
}

func (StoreCreated) Key() string {
	return StoreCreatedEvent
}
