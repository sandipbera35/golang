package main

import (
	"fmt"
)

type generics interface {
	int32 | int64 | float32 | float64 |
		int | uint32 | uint64 | string
}

func gener[T generics](v1 T, v2 T) T {

	var c T

	c = v1 + v2

	return c

}

func main() {
	// fmt.Println(gener([]int{1, 2, 3, 6, 5, 8}))
	// fmt.Println(gener([]string{"A ", "b ", "c ", "d "}))

	a := "Sandip "
	b := "Bera "

	c := 12
	d := 13

	e := 13.5
	q := 15.02

	f := gener(a, b)
	g := gener(c, d)
	h := gener(e, q)

	fmt.Println(f)
	fmt.Printf("Type is : %T \n", f)
	fmt.Println(g)
	fmt.Printf("Type is  : %T \n", g)
	fmt.Println(h)
	fmt.Printf("Type is  : %T \n", h)
}
