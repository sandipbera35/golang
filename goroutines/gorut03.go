package main

import (
	"fmt"
)

func main() {

	var data int

	go func() { data++ }()
	// time.Sleep(2 * time.Second)
	if data == 0 {
		fmt.Println("the value is 0.")
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
}
