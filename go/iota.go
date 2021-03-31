package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type WrappedBizError struct {
	wrapped error
	Code    uint
	Msg     string
}

func (e *WrappedBizError) Error() string {
	return errors.Wrap(e.wrapped, e.Msg).Error()
}

//go:generate stringer -type=BizError -output=iota_stringer.go
type BizError uint

func (e BizError) Error() string {
	return e.String()[3:]
}

func (e BizError) Newf(template string, args ...interface{}) error {
	return &WrappedBizError{
		wrapped: errors.Errorf(template, args...),
		Code:    uint(e),
		Msg:     e.Error(),
	}
}

func (e BizError) Wrap(err error) error {
	return &WrappedBizError{
		wrapped: err,
		Code:    uint(e),
		Msg:     e.Error(),
	}
}

func (e BizError) Wrapf(err error, template string, args ...interface{}) error {
	return &WrappedBizError{
		wrapped: errors.Wrapf(err, template, args...),
		Code:    uint(e),
		Msg:     e.Error(),
	}
}

const base = 10000

const (
	ErrTypeSystem = 100 + iota
	ErrTypeAPISetting
	ErrTypeBootstrap
	ErrTypeHeartbeat
	ErrTypeLogin
	ErrTypeUser
)

const (
	ErrMarshalFailed BizError = ErrTypeSystem*base + 1 + iota
	ErrUnmarshalFailed

	ErrUserNotFound BizError = ErrTypeUser*base + 1 + iota
	ErrUserUpdateNameFailed
	Err更新用户手机号码失败
)

// output:
// 1000001 => Marshal
// 1000002 => Unmarshal
// 1050003 => UserNotFound
// 1050004 => UserUpdateName
// 1050005 => 更新用户手机号码失败
func main() {
	for _, err := range []error{
		ErrMarshalFailed,
		ErrUnmarshalFailed,
		ErrUserNotFound,
		ErrUserUpdateNameFailed,
		Err更新用户手机号码失败,
	} {
		fmt.Printf("%d => %s\n", err, err)
	}

	fmt.Println(ErrUserUpdateNameFailed.Newf("%s", "ctx"))
	fmt.Println(ErrUserUpdateNameFailed.Wrap(errors.New("cause")))
	fmt.Println(ErrUserUpdateNameFailed.Wrapf(errors.New("cause"), "%s", "ctx"))
}
