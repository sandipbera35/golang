package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)

	}
}
func main() {
	f("direct")

	go f("Goroutine")

	go func(msg string) {

		fmt.Println(msg)

	}("Going")
	time.Sleep(2 * time.Second)
	fmt.Println("Done!")
}
