package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println(a)
	a[4] = 100

	fmt.Println("set", a)
	fmt.Println("get", a[4])

	var twoD [5][6]int

	for i := 0; i < 5; i++ {
		for j := 0; j < 6; j++ {

			twoD[i][j] = i + j

		}
	}
	//fmt.Printf("2d: %v \n", twoD)

	for k, i := range twoD {

		fmt.Printf("index = %v  value = %v \n", k, i)

	}

}
