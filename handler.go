package bus

import (
	"context"
)

type Handler[T any] interface {
	Execute(ctx context.Context, data T) error
}

type HandlerFunc[T any] func(ctx context.Context, data T) error

func (fn HandlerFunc[T]) Execute(ctx context.Context, data T) error {
	return fn(ctx, data)
}

func Subsrcibe[T any](ctx context.Context, bus *Bus, event string, handler Handler[T]) {
	bus.On(ctx, event, NewSubscriber(func(ctx context.Context, event Event) {
		if data, ok := event.Message.(T); ok {
			err := handler.Execute(ctx, data)
			if err != nil {
				bus.logger.WithError(err).Errorf(ctx, "Execute")
			}
		}
	}))
}
