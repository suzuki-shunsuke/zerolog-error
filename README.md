# zerolog-error

[![Go Reference](https://pkg.go.dev/badge/github.com/suzuki-shunsuke/zerolog-error.svg)](https://pkg.go.dev/github.com/suzuki-shunsuke/zerolog-error)

Embed fields into error for [zerolog](https://github.com/rs/zerolog).

Go small library to embed fields into error for zerolog.

## Similar Projects

- [logrus-error](https://github.com/suzuki-shunsuke/logrus-error): Embed [logrus](https://github.com/sirupsen/logrus).Fields into error
- [zap-error](https://github.com/suzuki-shunsuke/zap-error): Embed [zap](https://github.com/uber-go/zap).Field into error

## Motivation

Embed structured data into error and output it with zerolog.

[fmt.Errorf](https://pkg.go.dev/fmt#Errorf) enables to add additional context to error.

e.g.

```go
fmt.Errorf("get a user: %w", err)
```

[rs/zerolog](https://github.com/rs/zerolog) is one of most popular structured logging library.

e.g.

```go
logger.Error().Err(err).Str("username", username).Msg("get a user")
```

`fmt.Errorf` is very useful, but you can add only a string to error as context. You can't add structured data to error.
If you use zerolog, you may want to add structured data to error.

This library provides small APIs to embed structured data into error and output it with zerolog.

## How to use

```go
import (
	"github.com/suzuki-shunsuke/zerolog-error/zerr"
)
```

There are two APIs.

- `WithFields(err error, fields ...Field) error`: Embed fields into error
- `WithError(ev *zerolog.Event, err error) *zerolog.Event`: Get fields from error and add them to an event

e.g.

```go
func updateUser() error {
	// ...
	return zerr.WithFields(errors.New("get a user"), &zerr.Str{"id", "foo"})
}
```

```go
	if err := updateUser(); err != nil {
		zerr.WithError(log.Error(), err).Send("update a user")
	}
```

## Document

Please see https://pkg.go.dev/github.com/suzuki-shunsuke/zerolog-error/zerr

## LICENSE

[MIT](LICENSE)
