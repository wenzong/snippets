package main

import "fmt"

/* Message */
type Message string

func NewMessage() Message {
	return Message("Hi there!")
}

/* Greeter */
type Greeter struct {
	Message Message
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewGreeter(m Message) *Greeter {
	return &Greeter{Message: m}
}

/* Event */
type Event struct {
	Greeter *Greeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func NewEvent(g *Greeter) Event {
	return Event{Greeter: g}
}

/* main */
func main() {
	e := InitEvent()

	e.Start()
}
