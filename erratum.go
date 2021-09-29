package erratum

import (
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/cockroachdb/errors/markers"
)

// This file demonstrates how to add a wrapper type not otherwise
// known to the rest of the library.

// WithFields is our wrapper type.
type withFields struct {
	cause error
	fields  Fields
}

type Fields map[string]interface{}

// WrapWithFields adds a HTTP code to an existing error.
func WrapWithFields(err error, fields Fields) error {
	if err == nil {
		return nil
	}
	return &withFields{cause: err, fields: fields}
}

// GetFields retrieves the HTTP code from a stack of causes.
func GetFields(err error) Fields {
	if v, ok := markers.If(err, func(err error) (interface{}, bool) {
		if w, ok := err.(*withFields); ok {
			return w.fields, true
		}
		return nil, false
	}); ok {
		return v.(Fields)
	}
	return nil
}

// it's an error.
func (w *withFields) Error() string { return w.cause.Error() }

// Cause makes it also a wrapper.
func (w *withFields) Cause() error  { return w.cause }
func (w *withFields) Unwrap() error { return w.cause }

// Format knows how to format itself.
func (w *withFields) Format(s fmt.State, verb rune) { errors.FormatError(w, s, verb) }

// SafeFormatError implements errors.SafeFormatter.
// Note: see the documentat ion of errbase.SafeFormatter for details
// on how to implement this. In particular beware of not emitting
// unsafe strings.
func (w *withFields) SafeFormatError(p errors.Printer) (next error) {
	if p.Detail() {
		p.Printf("fields: %+v", w.fields)
	}
	return w.cause
}
