package zerr

import (
	"errors"
	"reflect"
	"testing"
)

func Test_zError_Error(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		err  *zError
		exp  string
	}{
		{
			name: "nil",
		},
		{
			name: "embedded error is nil",
			err:  &zError{},
		},
		{
			name: "normal",
			err: &zError{
				err: errors.New("foo"),
				fields: []Field{
					Str("name", "yoo"),
				},
			},
			exp: "foo",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			m := d.err.Error()
			if d.exp != m {
				t.Fatalf("wanted %s, got %s", d.exp, m)
			}
		})
	}
}

func Test_zError_Unwrap(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		err  *zError
		exp  error
	}{
		{
			name: "nil",
		},
		{
			name: "embedded error is nil",
			err:  &zError{},
		},
		{
			name: "normal",
			err: &zError{
				err: errors.New("foo"),
				fields: []Field{
					Str("name", "yoo"),
				},
			},
			exp: errors.New("foo"),
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			m := d.err.Unwrap()
			if !reflect.DeepEqual(d.exp, m) {
				t.Fatalf("wanted %v, got %v", d.exp, m)
			}
		})
	}
}

func Test_zError_Fields(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		err  *zError
		exp  []Field
	}{
		{
			name: "nil",
		},
		{
			name: "no field",
			err: &zError{
				err: errors.New("foo"),
			},
		},
		{
			name: "normal",
			err: &zError{
				err: errors.New("foo"),
				fields: []Field{
					Str("name", "yoo"),
				},
			},
			exp: []Field{
				Str("name", "yoo"),
			},
		},
		{
			name: "nested",
			err: &zError{
				err: &zError{
					err: errors.New("foo"),
					fields: []Field{
						Str("age", "10"),
					},
				},
				fields: []Field{
					Str("name", "yoo"),
				},
			},
			exp: []Field{
				Str("age", "10"),
				Str("name", "yoo"),
			},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			m := d.err.Fields()
			if !reflect.DeepEqual(d.exp, m) {
				t.Fatalf("wanted %v, got %v", d.exp, m)
			}
		})
	}
}
