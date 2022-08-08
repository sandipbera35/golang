package main

import "fmt"

// creating interace
type generic interface {
	area() float64
}

// circle type declare
type circle struct {
	redious float64
}

// rect type created
type rect struct {
	width, height float64
}

// creating methods thats return area o circle
func (c circle) area() float64 {

	return 3.14 * c.redious * c.redious
}

func (r rect) area() float64 {
	return r.height * r.width
}
func calculate(g generic) {
	fmt.Println(g)
	fmt.Println(g.area())

}
func main() {

	r := rect{width: 3, height: 4}
	c := circle{redious: 5}
	calculate(r)
	calculate(c)
}
