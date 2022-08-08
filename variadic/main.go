package main

import "fmt"

func vari(ver ...int) {

	for _, i := range ver {

		fmt.Println(i)
	}

}

func main() {

	a := []int{1, 2, 3, 9, 6, 6, 4, 5, 6, 7, 8}

	vari(a...)

}
