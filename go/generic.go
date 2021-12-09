package main

import (
	"fmt"
)

func print[T any](s []T) {
	for _, v := range s{
		fmt.Println(v)
	}
}

// go run -gcflags=-G=3 generic.go
func main() {
	print([]string{"hello", "world"})
}
