package bus

import (
	"context"
	"encoding/json"
	"github.com/adverax/log"
)

// Sniffer is interface for sniffing events
type Sniffer interface {
	Asserted(ctx context.Context, subject string, subscriber Subscriber)
	Retracted(ctx context.Context, subject string, subscriber Subscriber)
	Publish(ctx context.Context, notification Notification)
}

type publisherSniffer struct {
	logger log.Logger
}

func (that *publisherSniffer) Asserted(ctx context.Context, subject string, subscriber Subscriber) {
	that.logger.Debugf(ctx, "BUS SUBSCRIBER ASSERTED: %s", subject)
}

func (that *publisherSniffer) Retracted(ctx context.Context, subject string, subscriber Subscriber) {
	that.logger.Debugf(ctx, "BUS SUBSCRIBER RETRACTED: %s", subject)
}

func (that *publisherSniffer) Publish(ctx context.Context, notification Notification) {
	fields := log.Fields{
		log.FieldKeyEntity:  "BUS",
		log.FieldKeyAction:  "<<",
		log.FieldKeySubject: notification.Subject,
	}

	if notification.Message != nil {
		data, _ := json.Marshal(notification.Message)
		fields[log.FieldKeyData] = string(data)
	}

	that.logger.
		WithFields(fields).
		Debug(ctx, "Event")
}

// NewSniffer is constructor for creating instance of the publisher sniffer
func NewSniffer(logger log.Logger) Sniffer {
	return &publisherSniffer{logger: logger}
}

type dummyPublisherSniffer struct{}

func (that *dummyPublisherSniffer) Asserted(ctx context.Context, subject string, subscriber Subscriber) {
}

func (that *dummyPublisherSniffer) Retracted(ctx context.Context, subject string, subscriber Subscriber) {
}

func (that *dummyPublisherSniffer) Publish(ctx context.Context, notification Notification) {}
