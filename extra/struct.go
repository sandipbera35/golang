package main

import "fmt"

type stru struct {
	name string
	age  int
}

func main() {

	s := stru{
		name: "Nandalal Bera",
		age:  58,
	}
	fmt.Println(s)
	fmt.Println(s.age)
	fmt.Println(s.name)

	sp := &s

	fmt.Println("sp:", sp.age)
	fmt.Println("sp ptr:", &sp)

}
