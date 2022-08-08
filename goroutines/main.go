package main

import (
	"fmt"
	"time"
)

func msg(msg string) {

	for i := 0; i < 5; i++ {
		fmt.Println(msg, "=", i)
		time.Sleep(time.Second)
	}
}
func msg2(msg string) {

	for i := 0; i < 5; i++ {
		fmt.Println(msg, "=", i)
		time.Sleep(time.Second)
	}
}
func msg3(msg string) {

	for i := 0; i < 5; i++ {
		fmt.Println(msg, "=", i)
		time.Sleep(time.Second)
	}
}
func main() {
	// go routine / green thread
	// go threads are not managed by os
	msg("No thread")
	go msg2("Go routine1")
	go msg3("Go routine2")

	fmt.Scanln()

}
