package main

import (
	"fmt"
)

//go:generate stringer -type=BizError -output=iota_stringer.go
type BizError uint

func (e BizError) Error() string {
	return e.String()[3:]
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
}
