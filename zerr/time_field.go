package zerr

import (
	"time"

	"github.com/rs/zerolog"
)

type timeField struct {
	Key   string
	Value time.Time
}

func (field timeField) With(ev *zerolog.Event) *zerolog.Event {
	return ev.Time(field.Key, field.Value)
}

func Time(key string, value time.Time) Field {
	return timeField{
		Key:   key,
		Value: value,
	}
}
