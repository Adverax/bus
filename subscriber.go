package bus

import (
	"github.com/adverax/events"
)

type Subscriber = events.Subscriber

type SubscriberFunc = events.SubscriberFunc

type subscribers []Subscriber

func (that subscribers) indexOf(action Subscriber) int {
	for i, act := range that {
		if action == act {
			return i
		}
	}
	return -1
}
