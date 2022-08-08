package main

import (
	"fmt"
	"time"
)

func dispaly(str string) {

	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)

	}
}

func main() {
	/* messages := make(chan string)

	go func() { messages <- "Sandip" }()

	msg := <-messages
	fmt.Println(msg) */

	go dispaly("Sandip")
	dispaly("Love GeeksforGeeks")
	time.Sleep(1 * time.Second)

}
