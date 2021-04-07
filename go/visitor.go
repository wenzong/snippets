package main

import "fmt"

// Elements
type Elem interface {
	Accept(Visitor)
}

type Head struct {
	head string
}

func (h *Head) Accept(v Visitor) {
	v(h)
}

type Body struct {
	body string
}

func (b *Body) Accept(v Visitor) {
	v(b)
}

// Visitor
type Visitor func(Elem)

func EchoType(e Elem) {
	fmt.Printf("%T\n", e)
}

func EchoValue(e Elem) {
	fmt.Printf("%s\n", e)
}

func main() {
	for _, e := range []Elem{&Head{head: "h1"}, &Body{body: "b"}} {
		for _, v := range []Visitor{EchoType, EchoValue} {
			e.Accept(v)
		}
	}
}
