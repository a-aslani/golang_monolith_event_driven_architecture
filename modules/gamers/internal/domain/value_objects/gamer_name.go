package value_objects

import (
	"github.com/stackus/errors"
)

var (
	ErrGamerLastNameIsBlank = errors.Wrap(errors.ErrBadRequest, "the gamer last name cannot be blank")
)

type GamerName struct {
	FistName string
	LastName string
}

func NewGamerName(firstName, lastName string) (GamerName, error) {

	if lastName == "" {
		return GamerName{}, ErrGamerLastNameIsBlank
	}

	return GamerName{
		FistName: firstName,
		LastName: lastName,
	}, nil
}
