package main

import (
	"fmt"
)

func foo(c chan int, someValue int) {
	c <- someValue * 5
}
func main() {

	ch := make(chan int)

	go foo(ch, 5)
	go foo(ch, 3)

	v1 := <-ch
	v2 := <-ch

	fmt.Println(v1, v2)

}
