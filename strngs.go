package main

import (
	"bytes"
	"fmt"
)

func main() {
	// for true {
	var l int

	fmt.Println("Enter lenth..")
	fmt.Scan(&l)

	strs := make([]string, l)

	for i := 0; i < l; i++ {
		var s string
		fmt.Printf("Enter strs[%v] : ", i)
		fmt.Scan(&s)
		// for i := 0; i < len(s); i++
		if len(s) == 5 {

			strs[i] = s

		} else {

			fmt.Printf("\n")
			fmt.Println("lenth of the word should be exactly 5 please reenter..")
			i = i - 1
		}
	}

	var str3 bytes.Buffer

	for i := 0; i < len(strs); i++ {

		var sb string

		sb = strs[i]

		if len(sb) != 5 {
			i--

		} else {

			str1 := sb[2:3]

			str3.WriteString(str1)

			//fmt.Println("str1: ", str1)
		}
	}

	fmt.Printf("\n")
	fmt.Println("Strs :", strs)
	fmt.Printf("\n")
	fmt.Println("after Concatinetd every 3 letter of every words: ", str3.String())
	fmt.Printf("\n")

}

// }
