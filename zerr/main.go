package zerr

import (
	"errors"

	"github.com/rs/zerolog"
)

func WithFields(err error, fields ...Field) error {
	if err == nil {
		return nil
	}
	return &zError{
		err:    err,
		fields: fields,
	}
}

func WithError(ev *zerolog.Event, err error) *zerolog.Event {
	if err == nil {
		return ev
	}
	var e *zError
	if errors.As(err, &e) {
		for _, field := range e.Fields() {
			ev = field.With(ev)
		}
		return ev.Err(e.Unwrap())
	}
	return ev.Err(err)
}

func toFields(err error) []Field {
	if err == nil {
		return nil
	}
	var e *zError
	if errors.As(err, &e) {
		return e.Fields()
	}
	return nil
}
