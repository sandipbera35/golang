package main

import (
	"fmt"
	"strings"
)

type details struct {
	name string
	age  int
}

// func (data *details) Apppend(name string, age int) {
// 	data.name = append(data.name, name)
// 	data.age = append(data.age, age)
// }

func main() {

	m := make(map[string]int)

	for true {

		// fmt.Println("mymap:", m)

		var ch int

		fmt.Println("Press 1 to enter details..")
		fmt.Println("Press 2 to show sata ..")
		fmt.Println("Press 3 to search...")
		fmt.Println("Press 4 to delete...")

		fmt.Scan(&ch)

		switch ch {

		case 1:

			// fmt.Println("enter the lenth of the map")
			// var l int
			// fmt.Scan(&l)
			//for i := 0; i < l; i++
			var s string
			var n int
			fmt.Println("Enter name :")
			fmt.Scan(&s)
			fmt.Println("Enter age :")
			fmt.Scan(&n)

			var a details

			a.name = s
			a.age = n

			m[a.name] = a.age

			// mCopy := make(map[string]int)
			// for k, v := range m {
			// 	mCopy[k] = v
			// }
			fmt.Println("map:", m)
		case 2:
			fmt.Println("Case 2")
			fmt.Printf("\n")
			for i, v := range m {

				fmt.Printf("name is %v age is %v \n", i, v)

			}
			fmt.Printf("\n")

		case 3:
			fmt.Println("Case 3")
			fmt.Println("Enter name to search..")
			var l string
			fmt.Scan(&l)

			for i, v := range m {

				if i != l {
					for i := range m {

						fmt.Println("User does not exist..")

						v := strings.Contains(i, l)
						if v != true {
							break

						} else {

							fmt.Println("did you mean :", i, " ?")
							// var ch string
							// fmt.Println("if yes enter Y if no enter N")
							break
						}

					}
				} else {

					fmt.Printf("age of %v is %v \n", i, v)
					continue
				}

			}

		case 4:
			fmt.Println("Delete block")

			fmt.Println("Enter name to delete.")
			var l string
			fmt.Scan(&l)
			for i := range m {

				if i == l {
					delete(m, i)
					fmt.Printf("%v deleted \n", l)
				}

				// fmt.Printf("i: %T \n", i)

			}

		default:
			fmt.Println("Pls enter valid number from listed..")

		}

	}
}
