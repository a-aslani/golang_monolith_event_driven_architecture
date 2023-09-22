package value_objects

import "github.com/stackus/errors"

var (
	ErrGamerGemInvalidValue = errors.Wrap(errors.ErrBadRequest, "the amount of gem is not valid")
)

type GamerGem struct {
	Value int
}

func NewGamerGem(value int) (GamerGem, error) {
	if value < 0 {
		return GamerGem{}, ErrGamerGemInvalidValue
	}

	return GamerGem{
		Value: value,
	}, nil
}
