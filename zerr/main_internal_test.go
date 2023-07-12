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

func BenchmarkWithFields(b *testing.B) {
	err := errors.New("foo")
	b.Run("normal", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			arr1 := make([]Field, 10)
			for i := 0; i < 10; i++ {
				arr1[i] = &Str{
					fmt.Sprintf("foo-%d", i),
					fmt.Sprintf("foo-%d", i),
				}
			}
			WithFields(err, arr1...) //nolint:errcheck
		}
	})
	b.Run("str2", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			arr2 := make([]Field, 10)
			for i := 0; i < 10; i++ {
				arr2[i] = String(
					fmt.Sprintf("foo-%d", i),
					fmt.Sprintf("foo-%d", i),
				)
			}
			WithFields(err, arr2...) //nolint:errcheck
		}
	})
}

func BenchmarkWithError(b *testing.B) {
	err := errors.New("foo")
	b.Run("normal", func(b *testing.B) {
		ev := log.Info()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			arr1 := make([]Field, 10)
			for i := 0; i < 10; i++ {
				arr1[i] = &Str{
					fmt.Sprintf("foo-%d", i),
					fmt.Sprintf("foo-%d", i),
				}
			}
			WithError(ev, WithFields(err, arr1...))
		}
	})
	b.Run("str2", func(b *testing.B) {
		ev := log.Info()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			arr2 := make([]Field, 10)
			for i := 0; i < 10; i++ {
				arr2[i] = String(
					fmt.Sprintf("foo-%d", i),
					fmt.Sprintf("foo-%d", i),
				)
			}
			WithError(ev, WithFields(err, arr2...))
		}
	})
}

type Str2 struct {
	Key   string
	Value string
}

func (field Str2) With(ev *zerolog.Event) *zerolog.Event {
	return ev.Str(field.Key, field.Value)
}

func String(key, value string) Str2 {
	return Str2{
		Key:   key,
		Value: value,
	}
}
