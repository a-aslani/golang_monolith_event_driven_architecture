package value_objects

import (
	"github.com/stackus/errors"
)

var (
	ErrGamerLastNameIsBlank = errors.Wrap(errors.ErrBadRequest, "the gamer last name cannot be blank")
)

type GamerName struct {
	fistName string
	lastName string
}

func NewGamerName(firstName, lastName string) (GamerName, error) {

	if lastName == "" {
		return GamerName{}, ErrGamerLastNameIsBlank
	}

	return GamerName{
		fistName: firstName,
		lastName: lastName,
	}, nil
}

func (g GamerName) FistName() string {
	return g.fistName
}

func (g GamerName) LastName() string {
	return g.lastName
}
