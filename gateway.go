package bus

import "context"

type MatchFilter interface {
	IsMatch(subject string) bool
}

type Gateway struct {
	publisher Publisher
	filter    MatchFilter
	prefix    string
	suffix    string
}

func (that *Gateway) Event(
	ctx context.Context,
	notification Notification,
) {
	if that.filter.IsMatch(notification.Subject) {
		notification.Subject = that.mute(notification.Subject)
		that.publisher.Publish(ctx, notification)
	}
}

func (that *Gateway) mute(subject string) string {
	return that.prefix + subject + that.suffix
}

// NewGateway is constructor for creating instance of gateway
func NewGateway(
	filter MatchFilter,
	publisher Publisher,
	prefix, suffix string,
) *Gateway {
	return &Gateway{
		publisher: publisher,
		filter:    filter,
		prefix:    prefix,
		suffix:    suffix,
	}
}
