package zerr

import "github.com/rs/zerolog"

type int64Field struct {
	Key   string
	Value int64
}

func (field int64Field) With(ev *zerolog.Event) *zerolog.Event {
	return ev.Int64(field.Key, field.Value)
}

func Int64(key string, value int64) Field {
	return int64Field{
		Key:   key,
		Value: value,
	}
}
