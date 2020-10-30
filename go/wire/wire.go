//+build wireinject

package main

import "github.com/google/wire"

var BaseSet = wire.NewSet(
	NewGreeter,
	NewMessage,
)

func InitEvent() Event {
	wire.Build(NewEvent, BaseSet)
	return Event{}
}
