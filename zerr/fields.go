package zerr

import "github.com/rs/zerolog"

type Field interface {
	With(ev *zerolog.Event) *zerolog.Event
}

type String struct {
	Key   string
	Value string
}

func (field *String) With(ev *zerolog.Event) *zerolog.Event {
	return ev.Str(field.Key, field.Value)
}
