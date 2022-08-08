package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 3

	fmt.Println("m:", m)

	j, prs := m["k1"]
	fmt.Println("prs:", prs)
	fmt.Println("j:", j)
}
