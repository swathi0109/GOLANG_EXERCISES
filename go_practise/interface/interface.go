package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float32
}

type Rectange struct {
	width  float32
	height float32
}

type shape interface {
	area() float32
}

func (r Rectange) area() float32 {
	area := r.width * r.height
	return area
}

func (c Circle) area() float32 {
	area1 := math.Pi * c.radius * c.radius
	return area1
}

func getarea(s shape) float32 {
	return s.area()
}

func main() {
	r1 := Rectange{5, 6}
	c1 := Circle{3}
	shapes := []shape{r1, c1}
	for _, shape := range shapes {
		fmt.Println(getarea(shape))
	}
	//fmt.Println(r1.area())
	//fmt.Println(c1.area())
}
