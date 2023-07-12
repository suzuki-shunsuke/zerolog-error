package zerr

import (
	"time"

	"github.com/rs/zerolog"
)

type durField struct {
	Key   string
	Value time.Duration
}

func (field durField) With(ev *zerolog.Event) *zerolog.Event {
	return ev.Dur(field.Key, field.Value)
}

func Dur(key string, value time.Duration) Field {
	return durField{
		Key:   key,
		Value: value,
	}
}
