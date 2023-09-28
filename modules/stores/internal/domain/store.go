package domain

import (
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/stores/internal/domain/value_objects"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/es"
	"github.com/stackus/errors"
)

const StoreAggregate = "stores.Store"

type Store struct {
	es.Aggregate
	name   value_objects.StoreName
	amount value_objects.StoreAmount
	price  value_objects.StorePrice
}

var _ interface {
	es.EventApplier
	es.Snapshotter
} = (*Store)(nil)

func NewStore(id string) *Store {
	return &Store{
		Aggregate: es.NewAggregate(id, StoreAggregate),
	}
}

func CreateStore(id string, name string, amount int, price float64) (*Store, error) {

	store := NewStore(id)

	storeAmount, err := value_objects.NewStoreAmount(amount)
	if err != nil {
		return nil, err
	}

	storePrice, err := value_objects.NewStorePrice(price)
	if err != nil {
		return nil, err
	}

	storeName, err := value_objects.NewStoreName(name)
	if err != nil {
		return nil, err
	}

	store.name = storeName
	store.amount = storeAmount
	store.price = storePrice

	store.AddEvent(StoreCreatedEvent, &StoreCreated{
		Name:   storeName,
		Amount: storeAmount,
		Price:  storePrice,
	})

	return store, nil
}

func (s *Store) Edit(name string, amount int, price float64) error {

	storeName, err := value_objects.NewStoreName(name)
	if err != nil {
		return err
	}

	storeAmount, err := value_objects.NewStoreAmount(amount)
	if err != nil {
		return err
	}
	storePrice, err := value_objects.NewStorePrice(price)
	if err != nil {
		return err
	}

	s.AddEvent(StoreEditedEvent, &StoreEdited{
		Name:   storeName,
		Amount: storeAmount,
		Price:  storePrice,
	})

	return nil
}

func (s *Store) Remove() error {
	s.AddEvent(StoreRemovedEvent, &StoreRemoved{})
	return nil
}

func (Store) Key() string {
	return StoreAggregate
}

func (s *Store) ApplyEvent(event ddd.Event) error {

	switch payload := event.Payload().(type) {
	case *StoreCreated:
		s.name = payload.Name
		s.amount = payload.Amount
		s.price = payload.Price

	case *StoreEdited:
		s.name = payload.Name
		s.amount = payload.Amount
		s.price = payload.Price

	case *StoreRemoved:

	default:
		return errors.ErrInternal.Msgf("%T received the event %s with unexpected payload %T", s, event.EventName(), payload)

	}

	return nil
}

func (s *Store) ApplySnapshot(snapshot es.Snapshot) error {

	switch ss := snapshot.(type) {
	case *StoreV1:
		s.name = ss.Name
		s.amount = ss.Amount
		s.price = ss.Price

	default:
		return errors.ErrInternal.Msgf("%T received the unexpected snapshot %T", s, snapshot)
	}

	return nil
}

func (s *Store) ToSnapshot() es.Snapshot {
	return StoreV1{
		Name:   s.name,
		Amount: s.amount,
		Price:  s.price,
	}
}
