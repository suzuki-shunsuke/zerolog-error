package zerr

import "github.com/rs/zerolog"

type boolField struct {
	Key   string
	Value bool
}

func (field boolField) With(ev *zerolog.Event) *zerolog.Event {
	return ev.Bool(field.Key, field.Value)

}

func Bool(key string, value bool) Field {
	return boolField{
		Key:   key,
		Value: value,
	}
}
