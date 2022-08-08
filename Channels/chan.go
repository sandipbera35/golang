package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)

	go func() {
		fmt.Println("Series 1")
		time.Sleep(1 * time.Second)
		done <- true
		done <- false
	}()
	fmt.Println("series 2")
	fmt.Println(<-done)

}
