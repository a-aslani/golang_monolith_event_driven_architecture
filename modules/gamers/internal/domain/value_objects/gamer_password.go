package value_objects

import "github.com/stackus/errors"

var (
	ErrGamerPasswordIsBlank = errors.Wrap(errors.ErrBadRequest, "the password can't be empty")
	ErrGamerPasswordIsMin6  = errors.Wrap(errors.ErrBadRequest, "the minimum valid password is 6 character")
)

type GamerPassword struct {
	Value string
}

func NewGamerPassword(password string) (GamerPassword, error) {

	if password == "" {
		return GamerPassword{}, ErrGamerPasswordIsBlank
	}

	if len(password) < 6 {
		return GamerPassword{}, ErrGamerPasswordIsMin6
	}

	return GamerPassword{Value: password}, nil
}
