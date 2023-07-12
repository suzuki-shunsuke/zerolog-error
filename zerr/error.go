package zerr

type zError struct {
	err    error
	fields []Field
}

func (e *zError) Error() string {
	if e == nil || e.err == nil {
		return ""
	}
	return e.err.Error()
}

func (e *zError) Unwrap() error {
	if e == nil || e.err == nil {
		return nil
	}
	return e.err
}

func (e *zError) Fields() []Field {
	if e == nil {
		return nil
	}
	return append(toFields(e.err), e.fields...)
}
