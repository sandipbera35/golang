package main

import "fmt"

func tesT() func() int {

	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	var1 := tesT()

	fmt.Println("var1-1", var1())
	fmt.Println("var1-2", var1())

}
