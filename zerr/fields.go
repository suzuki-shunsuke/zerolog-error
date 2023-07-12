package zerr

import "github.com/rs/zerolog"

type Field interface {
	With(ev *zerolog.Event) *zerolog.Event
}

type str struct {
	Key   string
	Value string
}

func (field str) With(ev *zerolog.Event) *zerolog.Event {
	return ev.Str(field.Key, field.Value)
}

func Str(key, value string) Field {
	return str{
		Key:   key,
		Value: value,
	}
}
