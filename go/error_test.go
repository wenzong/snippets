package main

import (
	"errors"
	"fmt"
	"testing"

	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

// Error Errorf ErrorFrom generate MyError.
func Error(msg string) error {
	return &MyError{wrapped: errors.New(msg)}
}

func Errorf(format string, a ...interface{}) error {
	return Error(fmt.Sprintf(format, a...))
}

func ErrorFrom(err error) error {
	return &MyError{
		wrapped: err,
	}
}

// MyError
type MyError struct {
	wrapped error
}

func (e *MyError) Error() string {
	return e.wrapped.Error()
}

func (e *MyError) Unwrap() error {
	return e.wrapped
}

// errors.As
//
// useful in interceptor to test return error type
func TestErrorsAs(t *testing.T) {
	for _, tc := range []struct {
		err    error
		expect string
	}{
		{
			err:    Error("from Error"),
			expect: "from Error",
		},
		{
			err:    Errorf("from %s", "Errorf"),
			expect: "from Errorf",
		},
		{
			err:    ErrorFrom(errors.New("from ErrorFrom")),
			expect: "from ErrorFrom",
		},
	} {
		assert.EqualError(t, tc.err, tc.expect)

		var err *MyError
		assert.True(t, errors.As(tc.err, &err))
	}
}

// errors.Is
func TestErrorsIs(t *testing.T) {
	err := errors.New("my")

	assert.True(t, errors.Is(ErrorFrom(err), err))
}

// errors.Unwrap
func TestErrorsUnwrap(t *testing.T) {
	err := errors.New("my")

	assert.Equal(t, err, errors.Unwrap(ErrorFrom(err)))
}

// github.com/pkg/errors.{New,Wrap,Cause}
func TestPkgErrorsWrap(t *testing.T) {
	err := pkgerrors.New("my")

	assert.Equal(t, err, pkgerrors.Cause(pkgerrors.Wrap(err, "wrapped")))
}
