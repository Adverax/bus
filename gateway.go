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
	event Event,
) {
	if that.filter.IsMatch(event.Subject) {
		event.Subject = that.mute(event.Subject)
		that.publisher.Publish(ctx, event)
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
