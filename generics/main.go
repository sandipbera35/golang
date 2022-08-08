package main

import (
	"fmt"
)

type generics interface {
	int | int64 | float64 | string | float32
}

func addInts(list []int) int {
	sum := 0
	for _, item := range list {
		sum += item
	}
	return sum
}
func addFloats(list []float64) float64 {
	var sum float64
	for _, item := range list {
		sum += item
	}
	return sum
}

func gener[T generics](list []T) T {
	var sum T
	for _, item := range list {
		sum += item
	}
	return sum

}

func main() {

	fmt.Println("int:", addInts([]int{1, 3, 2, 5, 6}))
	fmt.Println("float", addFloats([]float64{01, 0.3, 2.3, 5.8}))
	fmt.Println("gener with float:", gener([]float64{2.2, 3.0, 6.0, 7.0}))
	fmt.Println("gener with int:", gener([]int{2, 3, 6, 7}))
}
