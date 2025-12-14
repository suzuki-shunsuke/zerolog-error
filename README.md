# zerolog-error

[![DeepWiki](https://img.shields.io/badge/DeepWiki-suzuki--shunsuke%2Fzerolog--error-blue.svg?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACwAAAAyCAYAAAAnWDnqAAAAAXNSR0IArs4c6QAAA05JREFUaEPtmUtyEzEQhtWTQyQLHNak2AB7ZnyXZMEjXMGeK/AIi+QuHrMnbChYY7MIh8g01fJoopFb0uhhEqqcbWTp06/uv1saEDv4O3n3dV60RfP947Mm9/SQc0ICFQgzfc4CYZoTPAswgSJCCUJUnAAoRHOAUOcATwbmVLWdGoH//PB8mnKqScAhsD0kYP3j/Yt5LPQe2KvcXmGvRHcDnpxfL2zOYJ1mFwrryWTz0advv1Ut4CJgf5uhDuDj5eUcAUoahrdY/56ebRWeraTjMt/00Sh3UDtjgHtQNHwcRGOC98BJEAEymycmYcWwOprTgcB6VZ5JK5TAJ+fXGLBm3FDAmn6oPPjR4rKCAoJCal2eAiQp2x0vxTPB3ALO2CRkwmDy5WohzBDwSEFKRwPbknEggCPB/imwrycgxX2NzoMCHhPkDwqYMr9tRcP5qNrMZHkVnOjRMWwLCcr8ohBVb1OMjxLwGCvjTikrsBOiA6fNyCrm8V1rP93iVPpwaE+gO0SsWmPiXB+jikdf6SizrT5qKasx5j8ABbHpFTx+vFXp9EnYQmLx02h1QTTrl6eDqxLnGjporxl3NL3agEvXdT0WmEost648sQOYAeJS9Q7bfUVoMGnjo4AZdUMQku50McDcMWcBPvr0SzbTAFDfvJqwLzgxwATnCgnp4wDl6Aa+Ax283gghmj+vj7feE2KBBRMW3FzOpLOADl0Isb5587h/U4gGvkt5v60Z1VLG8BhYjbzRwyQZemwAd6cCR5/XFWLYZRIMpX39AR0tjaGGiGzLVyhse5C9RKC6ai42ppWPKiBagOvaYk8lO7DajerabOZP46Lby5wKjw1HCRx7p9sVMOWGzb/vA1hwiWc6jm3MvQDTogQkiqIhJV0nBQBTU+3okKCFDy9WwferkHjtxib7t3xIUQtHxnIwtx4mpg26/HfwVNVDb4oI9RHmx5WGelRVlrtiw43zboCLaxv46AZeB3IlTkwouebTr1y2NjSpHz68WNFjHvupy3q8TFn3Hos2IAk4Ju5dCo8B3wP7VPr/FGaKiG+T+v+TQqIrOqMTL1VdWV1DdmcbO8KXBz6esmYWYKPwDL5b5FA1a0hwapHiom0r/cKaoqr+27/XcrS5UwSMbQAAAABJRU5ErkJggg==)](https://deepwiki.com/suzuki-shunsuke/zerolog-error)

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
	return zerr.WithFields(errors.New("get a user"), zerr.Str("id", "foo"))
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
