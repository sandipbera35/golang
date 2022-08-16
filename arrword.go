package main

import (
	"bytes"
	"fmt"
)

func main() {
	// start:
	for true {
		var l int
		fmt.Println("Enter lenth..")
		fmt.Scan(&l)

		strs := make([]string, l)

		for i := 0; i < l; i++ {

			var s string

			fmt.Printf("Enter String with 5 latter in strs[%v]:\n ", i)

			fmt.Scan(&s)

			if len(s) != 5 {
				fmt.Printf("\n\n")
				fmt.Printf("every word lenth should be exact 5 enter again... \n\n")

				break

			} else {
				strs[i] = s
			}

		}

		var s1 bytes.Buffer

		for i := 0; i < len(strs); i++ {

			var word string

			word = strs[i]

			lent := len(word)

			if lent != 5 {
				break
			} else {

				byt := word[2:3]

				s1.WriteString(byt)

			}

		}

		fmt.Printf("\n\n")

		fmt.Println("strs:", strs)

		fmt.Printf("\n\n")

		fmt.Println("concatinated string:", s1.String())

		fmt.Printf("\n\n")

	}

}
