package main

import (
	"fmt"
	"time"
)

func PrintU(a string) {

	for i := 0; i < len(a); i++ {
		fmt.Printf("%v \n", a)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	go PrintU("Sandip")

	//fmt.Println("")
	go PrintU("Bera")
	time.Sleep(2 * time.Second)

	go PrintU("Bimal")
	time.Sleep(2 * time.Second)

}
