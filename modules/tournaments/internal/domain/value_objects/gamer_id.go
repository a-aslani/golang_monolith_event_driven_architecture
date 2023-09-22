package value_objects

import (
	"github.com/stackus/errors"
	"strings"
)

var (
	ErrGamerIdIsRequired = errors.Wrap(errors.ErrBadRequest, "the gamer ID is required")
)

type GamerID struct {
	Value string
}

func NewGamerID(id string) (GamerID, error) {

	if strings.TrimSpace(id) == "" {
		return GamerID{}, ErrGamerIdIsRequired
	}

	return GamerID{
		Value: id,
	}, nil
}
