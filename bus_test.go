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
		SubscriberFunc(func(ctx context.Context, notification Notification) {
			fmt.Println("event:", notification.Subject, notification.Message)
		}),
	)
	bus.Publish(ctx, Notification{
		Subject: "subject",
		Message: "message",
	})
	time.Sleep(time.Second)

	// Output:
	// event: subject message
}
