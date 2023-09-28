package value_objects

import (
	"github.com/stackus/errors"
	"strings"
)

var (
	ErrStoreNameIsBlank = errors.Wrap(errors.ErrBadRequest, "the name field is required")
)

type StoreName struct {
	Value string
}

func NewStoreName(name string) (StoreName, error) {

	if len(strings.TrimSpace(name)) <= 0 {
		return StoreName{}, ErrStoreNameIsBlank
	}

	return StoreName{
		Value: name,
	}, nil
}
