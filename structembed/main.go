package main

import "fmt"

type person struct {
	name    string
	surname string
	id      int
}
type Licence struct {
	person
	ltk string
}

func main() {
	p1 := person{
		name:    "Sandip",
		surname: "Bera",
		id:      05,
	}
	p2 := Licence{
		person: person{
			name:    "James",
			surname: "Bond",
			id:      07,
		},
		ltk: "Yes",
	}
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)

}
