package zerr

import "github.com/rs/zerolog"

type Field interface {
	With(ev *zerolog.Event) *zerolog.Event
}
