package main

import (
	"fmt"
)

func main() {

	var arr = [8]int32{1, 2, 3, 4, 5, 65, 10, 88}

	for i := 0; i < 6; i++ {

		fmt.Println("Pos: ", i, " Value : ", arr[i], "Addr: ", &arr[i])
	}

}
