package bus

import (
	"encoding/json"
	"fmt"
	"github.com/adverax/types"
	"time"
)

type Notification struct {
	Subject string      `json:"subject"`           // Название темы
	Message interface{} `json:"message,omitempty"` // Тело сообщения
}

func (that *Notification) String() string {
	data, _ := json.Marshal(that.Message)
	return fmt.Sprintf("%s %s", that.Subject, string(data))
}

func (that *Notification) AsRawMessage(defaults json.RawMessage) json.RawMessage {
	return types.Type.Json.Cast(that.Message, defaults)
}

func (that *Notification) AsString(defaults string) string {
	return types.Type.String.Cast(that.Message, defaults)
}

func (that *Notification) AsInteger(defaults int64) int64 {
	return types.Type.Integer.Cast(that.Message, defaults)
}

func (that *Notification) AsFloat(defaults float64) float64 {
	return types.Type.Float.Cast(that.Message, defaults)
}

func (that *Notification) AsBoolean(defaults bool) bool {
	return types.Type.Boolean.Cast(that.Message, defaults)
}

func (that *Notification) AsDuration(defaults time.Duration) time.Duration {
	return types.Type.Duration.Cast(that.Message, defaults)
}
