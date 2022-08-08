package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area()
	perim()
}

type rect struct {
	width  float64
	height float64
}
type circle struct {
	redious float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.redious * c.redious
}

func main() {

	r := rect{width: 20, height: 30}
	c := circle{redious: 6}

	fmt.Println("r:", r)
	fmt.Println("c", c)

	fmt.Println(circle.area(c))
	fmt.Println(rect.area(r))

}
