package main

import (
	"fmt"
	"time"
)

func cll(val int) {

	for i := 0; i < val; i++ {

		time.Sleep(time.Second)
		fmt.Println(i)

	}

}

func main() {
	go cll(10)
	go cll(8)
	fmt.Scanln()
}


