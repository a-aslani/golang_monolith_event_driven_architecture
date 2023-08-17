package value_objects

import (
	"github.com/stackus/errors"
	"strings"
)

var (
	ErrTournamentNameIsRequired = errors.Wrap(errors.ErrBadRequest, "tournament name is required")
)

type TournamentName struct {
	value string
}

func NewTournamentName(name string) (TournamentName, error) {

	if strings.TrimSpace(name) == "" {
		return TournamentName{}, ErrTournamentNameIsRequired
	}

	return TournamentName{
		value: name,
	}, nil
}

func (receiver TournamentName) Value() string {
	return receiver.value
}
