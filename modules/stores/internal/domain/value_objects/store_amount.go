package value_objects

import "github.com/stackus/errors"

var (
	ErrStoreAmountInvalidValue = errors.Wrap(errors.ErrBadRequest, "the amount is invalid")
)

type StoreAmount struct {
	Value int
}

func NewStoreAmount(value int) (StoreAmount, error) {
	if value <= 0 {
		return StoreAmount{}, ErrStoreAmountInvalidValue
	}

	return StoreAmount{Value: value}, nil
}
