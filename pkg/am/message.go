package am

import (
	"context"
	"github.com/a-aslani/golang_monolith_event_driven_architecture/pkg/ddd"
)

type (
	Message interface {
		ddd.IDer
		MessageName() string
		Ack() error
		NAck() error
		Extend() error
		Kill() error
	}

	IncomingMessage interface {
		Message
		Ack() error
		NAck() error
		Extend() error
		Kill() error
	}

	MessageHandler[I Message] interface {
		HandleMessage(ctx context.Context, msg I) error
	}

	MessageHandlerFunc[I Message] func(ctx context.Context, msg I) error

	MessagePublisher[I any] interface {
		Publish(ctx context.Context, topicName string, v I) error
	}

	MessageSubscriber[O Message] interface {
		Subscribe(topicName string, handler MessageHandler[O], options ...SubscriberOption) error
	}

	MessageStream[I any, O Message] interface {
		MessagePublisher[I]
		MessageSubscriber[O]
	}
)

func (f MessageHandlerFunc[O]) HandleMessage(ctx context.Context, msg O) error {
	return f(ctx, msg)
}
