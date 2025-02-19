package bus

import (
	"encoding/json"
	"fmt"
	"github.com/adverax/types"
	"time"
)

type Event struct {
	Subject string      `json:"subject"`
	Message interface{} `json:"message,omitempty"`
}

func (that *Event) String() string {
	data, _ := json.Marshal(that.Message)
	return fmt.Sprintf("%s %s", that.Subject, string(data))
}

func (that *Event) AsRawMessage(defaults json.RawMessage) json.RawMessage {
	return types.Type.Json.Cast(that.Message, defaults)
}

func (that *Event) AsString(defaults string) string {
	return types.Type.String.Cast(that.Message, defaults)
}

func (that *Event) AsInteger(defaults int64) int64 {
	return types.Type.Integer.Cast(that.Message, defaults)
}

func (that *Event) AsFloat(defaults float64) float64 {
	return types.Type.Float.Cast(that.Message, defaults)
}

func (that *Event) AsBoolean(defaults bool) bool {
	return types.Type.Boolean.Cast(that.Message, defaults)
}

func (that *Event) AsDuration(defaults time.Duration) time.Duration {
	return types.Type.Duration.Cast(that.Message, defaults)
}
