package main

import (
	"fmt"
)

func main() {

	var x int

	x = 10

	var ptr *int = &x

	fmt.Println("Vlaue of X is : ", x)
	fmt.Println("variable ptr holds :", ptr)
	fmt.Println("value of ptr : ", *ptr)

}
