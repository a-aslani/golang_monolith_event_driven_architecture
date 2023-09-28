package value_objects

import "github.com/stackus/errors"

var (
	ErrStorePriceInvalidValue = errors.Wrap(errors.ErrBadRequest, "the price is invalid")
)

type StorePrice struct {
	Value float64
}

func NewStorePrice(value float64) (StorePrice, error) {

	if value <= 0 {
		return StorePrice{}, ErrStorePriceInvalidValue
	}

	return StorePrice{Value: value}, nil
}
