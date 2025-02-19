package bus

import (
	"context"
	"github.com/adverax/events"
	"github.com/adverax/log"
	"sync"
)

type Event = events.Event

// Registrar is interface of registrar event handlers.
type Registrar interface {
	On(ctx context.Context, subject string, subscriber Subscriber)
	Off(ctx context.Context, subject string, subscriber Subscriber)
}

// Publisher is interface for publishing events
type Publisher interface {
	Publish(ctx context.Context, event Event)
}

type Bus struct {
	sync.Mutex
	subscribers map[string]subscribers
	sniffer     Sniffer
	logger      log.Logger
}

func (that *Bus) Cleanup() {
	that.Lock()
	defer that.Unlock()

	that.subscribers = make(map[string]subscribers, 4)
}

func (that *Bus) On(
	ctx context.Context,
	subject string,
	action Subscriber,
) {
	that.Lock()
	defer that.Unlock()

	if actions, found := that.subscribers[subject]; found {
		that.subscribers[subject] = append(actions, action)
	} else {
		that.subscribers[subject] = subscribers{action}
	}

	that.sniffer.Asserted(ctx, subject, action)
}

func (that *Bus) Off(
	ctx context.Context,
	subject string,
	action Subscriber,
) {
	that.Lock()
	defer that.Unlock()

	if list, found := that.subscribers[subject]; found {
		index := list.indexOf(action)
		if index >= 0 {
			that.subscribers[subject] = append(list[:index], list[index+1:]...)
			that.sniffer.Retracted(ctx, subject, action)
		}
	}
}

func (that *Bus) Publish(
	ctx context.Context,
	event Event,
) {
	that.sniffer.Publish(ctx, event)

	ss := that.getSubscribers(event.Subject)
	for _, subscriber := range ss {
		go subscriber.HandleEvent(ctx, event)
	}
}

func (that *Bus) getSubscribers(subject string) subscribers {
	that.Lock()
	defer that.Unlock()

	ss := that.subscribers[subject]
	result := make(subscribers, len(ss))
	copy(result, ss)
	return result
}

// New is constructor for creating instance of the publisher
func New(sniffer Sniffer, logger log.Logger) *Bus {
	if sniffer == nil {
		sniffer = &dummyPublisherSniffer{}
	}

	if logger == nil {
		logger = log.NewDummyLogger()
	}

	return &Bus{
		subscribers: make(map[string]subscribers, 4),
		sniffer:     sniffer,
		logger:      logger,
	}
}
