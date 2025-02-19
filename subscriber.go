package bus

import (
	"context"
)

// Subscriber is abstract subscriber for the event
type Subscriber interface {
	Event(ctx context.Context, event Event)
}

type subscriber struct {
	action func(ctx context.Context, event Event)
}

func (that *subscriber) Event(ctx context.Context, event Event) {
	that.action(ctx, event)
}

func NewSubscriber(action func(ctx context.Context, event Event)) Subscriber {
	return &subscriber{
		action: action,
	}
}

type subscribers []Subscriber

func (that subscribers) indexOf(action Subscriber) int {
	for i, act := range that {
		if action == act {
			return i
		}
	}
	return -1
}
