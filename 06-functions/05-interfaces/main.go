// ● create a type SQUARE
// ● create a type CIRCLE
// ● attach a method to each that calculates AREA and returns it
// ○ circle area= π r 2
// ○ square area = L * W
// ● create a type SHAPE that defines an interface as anything that has the AREA method
// ● create a func INFO which takes type shape and then prints the area
// ● create a value of type square
// ● create a value of type circle
// ● use func info to print the area of square
// ● use func info to print the area of circle

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Square struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.length * s.width
}

func (c Circle) area() float64 {
	r := c.radius
	return math.Pi * r * r
}

func info(s Shape) float64 {
	return s.area()
}

func main() {
	c := Circle{7}
	s := Square{4, 5}

	fmt.Println(info(c))
	fmt.Println(info(s))
}
