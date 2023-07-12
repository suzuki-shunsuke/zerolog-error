package zerr

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestWithFields(t *testing.T) {
	t.Parallel()
	data := []struct {
		name   string
		err    error
		fields []Field
		exp    error
	}{
		{
			name: "nil",
			fields: []Field{
				&Str{"id", "foo"},
			},
		},
		{
			name: "normal",
			err:  errors.New("get a user"),
			fields: []Field{
				&Str{"id", "foo"},
			},
			exp: &zError{
				err: errors.New("get a user"),
				fields: []Field{
					&Str{"id", "foo"},
				},
			},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			err := WithFields(d.err, d.fields...)
			if !reflect.DeepEqual(d.exp, err) {
				t.Fatal("must be same")
			}
		})
	}
}

func TestWithError(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		ev   *zerolog.Event
		exp  *zerolog.Event
		err  error
	}{
		{
			name: "nil",
			ev:   log.Info(),
			exp:  log.Info(),
		},
		{
			name: "not zError",
			ev:   log.Info().Str("id", "yoo"),
			exp:  log.Info().Str("id", "yoo").Err(errors.New("foo")),
			err:  errors.New("foo"),
		},
		{
			name: "zError",
			ev:   log.Info().Str("id", "yoo"),
			exp:  log.Info().Str("id", "yoo").Str("name", "FOO").Err(errors.New("foo")),
			err: &zError{
				err: errors.New("foo"),
				fields: []Field{
					&Str{"name", "FOO"},
				},
			},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			ev := WithError(d.ev, d.err)
			if !reflect.DeepEqual(d.exp, ev) {
				t.Fatal("must be same")
			}
		})
	}
}

func Test_toFields(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		exp  []Field
		err  error
	}{
		{
			name: "nil",
		},
		{
			name: "not zError",
			err:  errors.New("foo"),
		},
		{
			name: "zError",
			exp: []Field{
				&Str{"name", "FOO"},
			},
			err: fmt.Errorf("get a user: %w", &zError{
				err: errors.New("foo"),
				fields: []Field{
					&Str{"name", "FOO"},
				},
			}),
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			fields := toFields(d.err)
			if !reflect.DeepEqual(d.exp, fields) {
				t.Fatal("must be same")
			}
		})
	}
}
