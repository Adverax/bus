package bus

import (
	"context"
	"fmt"
)

// Subscriber is abstract subscriber for the event
type Subscriber interface {
	Event(ctx context.Context, notification Notification)
}

type SubscriberFunc func(ctx context.Context, notification Notification)

func (fn SubscriberFunc) Event(ctx context.Context, notification Notification) {
	fn(ctx, notification)
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
