package es

import (
	"fmt"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/registry"
)

type VersionSetter interface {
	setVersion(int64)
}

func SetVersion(version int64) registry.BuildOption {
	return func(v interface{}) error {
		if agg, ok := v.(VersionSetter); ok {
			agg.setVersion(version)
			return nil
		}
		return fmt.Errorf("%T does not have the method setVersion(int)", v)
	}
}
