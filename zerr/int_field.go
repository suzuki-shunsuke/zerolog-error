package zerr

import "github.com/rs/zerolog"

type intField struct {
	Key   string
	Value int
}

func (field intField) With(ev *zerolog.Event) *zerolog.Event {
	return ev.Int(field.Key, field.Value)
}

func Int(key string, value int) Field {
	return intField{
		Key:   key,
		Value: value,
	}
}
