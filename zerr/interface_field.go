package zerr

import "github.com/rs/zerolog"

type interfaceField struct {
	Key   string
	Value interface{}
}

func (field interfaceField) With(ev *zerolog.Event) *zerolog.Event {
	return ev.Interface(field.Key, field.Value)
}

func Interface(key string, value interface{}) Field {
	return interfaceField{
		Key:   key,
		Value: value,
	}
}
