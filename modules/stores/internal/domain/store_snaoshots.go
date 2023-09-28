package domain

import "github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/domain/value_objects"

type StoreV1 struct {
	Name   value_objects.StoreName
	Amount value_objects.StoreAmount
	Price  value_objects.StorePrice
}

func (s StoreV1) SnapshotName() string {
	return "stores.StoreV1"
}
