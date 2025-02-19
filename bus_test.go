package bus

import (
	"context"
	"fmt"
	"time"
)

func ExampleBus() {
	ctx := context.Background()
	bus := New(nil, nil)
	bus.On(
		ctx,
		"subject",
		SubscriberFunc(func(ctx context.Context, event Event) {
			fmt.Println("event:", event.Subject, event.Message)
		}),
	)
	bus.Publish(ctx, Event{
		Subject: "subject",
		Message: "message",
	})
	time.Sleep(time.Second)

	// Output:
	// event: subject message
}
