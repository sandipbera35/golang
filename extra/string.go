package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "สวัสดี"

	fmt.Println("len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x \n", s[i])
	}
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

}
