package main

import (
	"context"
	"fmt"
)

// Object which we want to change its behavior by change its opts field
type Object struct {
	opts options
}

// Make it public will allow user pass customize ObjectOption function
type options struct {
	height int
	width  int

	someBoolValue bool

	// could define as `type XXXInterceptor func(...) (...)`
	// someInterceptor XXXInterceptor
	someInterceptor func(ctx context.Context, req interface{}) (resp interface{}, err error)
}

var defaultOptions = options{
	height:          100,
	width:           100,
	someBoolValue:   true,
	someInterceptor: nil,
}

type ObjectOption func(*options)

// functional option 1
func Height(h int) ObjectOption {
	return func(o *options) {
		o.height = h
	}
}

// functional option 2
func SomeBoolValue(b bool) ObjectOption {
	return func(o *options) {
		o.someBoolValue = b
	}
}

// functional option 3
func SomeInterceptor(f func(ctx context.Context, req interface{}) (resp interface{}, err error)) ObjectOption {
	return func(o *options) {
		o.someInterceptor = f
	}
}

// NewObject is a [Variadic Function](https://en.wikipedia.org/wiki/Variadic_function)
// Accept ObjectOption...
func NewObject(opt ...ObjectOption) *Object {
	opts := defaultOptions

	for _, o := range opt {
		o(&opts)
	}

	return &Object{opts: opts}
}

func (object *Object) InspectOptions() {
	fmt.Println(object.opts.height)
	fmt.Println(object.opts.width)
	fmt.Println(object.opts.someBoolValue)
	fmt.Println(object.opts.someInterceptor)
}

func main() {
	object := NewObject(
		Height(1000),
		func(w int) ObjectOption {
			return func(o *options) {
				o.width = w
			}
		}(2000),
	)

	object.InspectOptions()
}
