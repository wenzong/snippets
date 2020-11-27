package main

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
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

var ErrMark = errors.New("mark error as MyError")

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

func (e *MyError) Is(target error) bool {
	return target == ErrMark
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

		assert.True(t, errors.Is(tc.err, ErrMark))
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

// errors.{Wrap,Cause}
func TesterrorsWrap(t *testing.T) {
	err := errors.New("my")

	assert.Equal(t, err, errors.Cause(errors.Wrap(err, "wrapped")))
}
