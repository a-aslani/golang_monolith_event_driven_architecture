package value_objects

import (
	"github.com/stackus/errors"
	"strings"
)

var (
	ErrTournamentDescriptionIsRequired = errors.Wrap(errors.ErrBadRequest, "description is required")
	ErrTournamentDescriptionMinLen     = errors.Wrap(errors.ErrBadRequest, "the length of the description must be greater than 5")
)

type TournamentDescription struct {
	Value string
}

func NewTournamentDescription(description string) (TournamentDescription, error) {

	if strings.TrimSpace(description) == "" {
		return TournamentDescription{}, ErrTournamentDescriptionIsRequired
	}

	if len(strings.TrimSpace(description)) < 5 {
		return TournamentDescription{}, ErrTournamentDescriptionMinLen
	}

	return TournamentDescription{
		Value: description,
	}, nil
}
