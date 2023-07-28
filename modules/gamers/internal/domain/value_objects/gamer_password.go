package value_objects

import "github.com/stackus/errors"

var (
	ErrGamerPasswordIsBlank = errors.Wrap(errors.ErrBadRequest, "the password can't be empty")
	ErrGamerPasswordIsMin6  = errors.Wrap(errors.ErrBadRequest, "the minimum valid password is 6 character")
)

type GamerPassword struct {
	value string
}

func NewGamerPassword(password string) (GamerPassword, error) {

	if password == "" {
		return GamerPassword{}, ErrGamerPasswordIsBlank
	}

	if len(password) < 6 {
		return GamerPassword{}, ErrGamerPasswordIsMin6
	}

	return GamerPassword{value: password}, nil
}

func (g GamerPassword) Value() string {
	return g.value
}
