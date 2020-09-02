package main

import "fmt"

func main() {
	// defer
	defer fmt.Println("world")

	// for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	defer fmt.Println("hello")

	sum := 1
	for sum < 1000 {
		sum += sum
	}

	for {
		break
	}

	// pointer/ptr
	i := 42
	p := &i
	fmt.Println(*p)

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.Y)

	p_v := &v
	p_v.X = 1e9
	// (*p_v).X = 1e9
	fmt.Println(v)
	fmt.Println(p_v)

	// arrays

	// slices

	// maps

	// make

}


// struct
type Vertex struct {
	X int
	Y int
}


// methods on types
func (v Vertex) Abs() float64 {
	return 1.0
}

// Interfaces

// goroutines

// channels

// sync.Mutex
