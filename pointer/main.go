package main

import (
	"fmt"
)

func main() {

	var x int

	x = 10

	fmt.Println("value of X: ", x)

	fmt.Println("Pointer addr ref of x : ", &x)

	var ptr *int = &x

	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr * 2

	fmt.Println(*ptr)
	fmt.Println(x)

}
