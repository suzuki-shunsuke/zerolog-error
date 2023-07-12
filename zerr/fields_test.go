package zerr_test

import (
	"reflect"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/suzuki-shunsuke/zerolog-error/zerr"
)

func TestStr_With(t *testing.T) {
	t.Parallel()
	data := []struct {
		name  string
		field *zerr.Str
		ev    *zerolog.Event
		exp   *zerolog.Event
	}{
		{
			name:  "normal",
			field: &zerr.Str{"id", "foo"},
			ev:    log.Info(),
			exp:   log.Info().Str("id", "foo"),
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			ev := d.field.With(d.ev)
			if !reflect.DeepEqual(d.exp, ev) {
				t.Fatal("must be same")
			}
		})
	}
}
