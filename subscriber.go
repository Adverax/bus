package bus

import (
	"context"
	"fmt"
)

// Subscriber is abstract subscriber for the event
type Subscriber interface {
	Event(ctx context.Context, event Event)
}

type SubscriberFunc func(ctx context.Context, event Event)

func (fn SubscriberFunc) Event(ctx context.Context, event Event) {
	fn(ctx, event)
}

type subscribers []Subscriber

func (that subscribers) indexOf(action Subscriber) int {
	for i, act := range that {
		if isEqualSubscribers(act, action) {
			return i
		}
	}
	return -1
}

func isEqualSubscribers(a, b Subscriber) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}
