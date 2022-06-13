package main

import (
	"fmt"
)

func print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	print([]string{"hello", "world"})

	fmt.Println(Sum([]int64{1, 2, 3, 4}))

	fmt.Println(Sum([]float64{0.1, 0.2, 1.1, 1.2}))

	fmt.Println(SumNumber([]int64{1, 2, 3, 4}))

	fmt.Println(SumNumber([]float64{0.1, 0.2, 1.1, 1.2}))
}

type Number interface {
	int | int64 | float64
}

func Sum[V int64 | float64](l []V) V {
	var sum V
	for _, v := range l {
		sum += v
	}

	return sum
}

func SumNumber[V Number](l []V) V {
	var sum V
	for _, v := range l {
		sum += v
	}

	return sum
}
