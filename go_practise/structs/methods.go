package main

import (
	"fmt"
	"math"
)

type Employee struct {
	f_name string
	l_name string
}

// using methods
func (e *Employee) changename(name string) {
	e.f_name = name
}

type Point struct {
	x int32
	y int32
}

type Circle struct {
	radius float32
	centre Point
}

func (c *Circle) setradius(rad float32) float32 {
	(*c).radius = rad
	return (*c).radius
}

func calc_area(c *Circle) float32 {
	area := math.Pi * c.radius * c.radius
	return area
}

func (p Point) changevalues(a, b int) (int32, int32) {
	p.x = int32(a)
	p.y = int32(b)
	return p.x, p.y
}

func main() {
	emp := Employee{"swathi", "shekar"}
	fmt.Println(emp)
	emp.changename("spoorthi")
	fmt.Println(emp)
	fmt.Println()

	p1 := Point{2, 4}
	c1 := Circle{3, p1}
	fmt.Println(c1)
	fmt.Println()
	area := float32(calc_area(&c1))
	fmt.Println(area)
	fmt.Println()
	c1.setradius(6)
	area2 := float32(calc_area(&c1))
	fmt.Println()
	fmt.Println(area2)
	fmt.Println()
	fmt.Println(c1.centre.x, c1.centre.y)
	x1, y1 := p1.changevalues(5, 6)
	fmt.Printf("After changing x=%d, y=%d", x1, y1)
}
