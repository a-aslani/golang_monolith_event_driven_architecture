package value_objects

import (
	"github.com/stackus/errors"
	"regexp"
)

var (
	ErrGamerEmailIsBlank   = errors.Wrap(errors.ErrBadRequest, "the email cannot be empty")
	ErrGamerEmailIsInvalid = errors.Wrap(errors.ErrBadRequest, "the email address is invalid type")
)

type GamerEmail struct {
	Value string
}

func NewGamerEmail(email string) (GamerEmail, error) {

	if email == "" {
		return GamerEmail{}, ErrGamerEmailIsBlank
	}

	if !isEmailValid(email) {
		return GamerEmail{}, ErrGamerEmailIsInvalid
	}

	return GamerEmail{
		Value: email,
	}, nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
