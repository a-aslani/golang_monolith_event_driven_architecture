package es

import (
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
)

type (
	Versioner interface {
		Version() int64
		PendingVersion() int64
	}

	Aggregate struct {
		ddd.Aggregate
		version int64
	}
)

var _ interface {
	EventCommitter
	Versioner
	VersionSetter
} = (*Aggregate)(nil)

func NewAggregate(id, name string) Aggregate {
	return Aggregate{
		Aggregate: ddd.NewAggregate(id, name),
	}
}

func (a *Aggregate) AddEvent(name string, data interface{}, options ...ddd.EventOption) {
	options = append(
		options,
		ddd.Metadata{
			ddd.AggregateVersionKey: a.PendingVersion() + 1,
		},
	)

	a.Aggregate.AddEvent(name, data, options...)
}

func (a *Aggregate) CommitEvents() {
	a.version += int64(len(a.Events()))
	a.ClearEvents()
}

func (a Aggregate) Version() int64            { return a.version }
func (a Aggregate) PendingVersion() int64     { return a.version + int64(len(a.Events())) }
func (a *Aggregate) setVersion(version int64) { a.version = version }
