package zerr_test

import (
	"errors"

	"github.com/rs/zerolog/log"
	"github.com/suzuki-shunsuke/zerolog-error/zerr"
)

func Example() {
	logger := log.With().Str("program", "example").Logger()
	if err := updateUser(); err != nil {
		// WithError gets an error and fields from an error
		// and returns *zerolog.Event
		zerr.WithError(logger.Error(), err).
			Str("type", "test").Msg("update a user")
	}
}

func updateUser() error {
	// WithFields adds fields to error
	return zerr.WithFields(errors.New("get a user"), zerr.Str("id", "foo"), zerr.Str("name", "Foo"))
}
