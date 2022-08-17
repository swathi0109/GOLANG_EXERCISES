package main

import (
	"fmt"
	"math"
)

func main() {
	/*var n float64
	fmt.Println("Enter the number")
	fmt.Scanln(&n)

	//square root
	result := math.Sqrt(n)
	fmt.Println("Square oot of a number is ", result)*/

	//quadratic equation
	var a, b, c, d float64
	fmt.Println("Enter the values for a, b, c")
	fmt.Scanln(&a, &b, &c)
	d = float64((b * b) - 4*a*c)
	fmt.Println(d)
	if d > 0 {
		r1 := (-b + math.Sqrt(d)) / (2 * a)
		r2 := (-b - math.Sqrt(d)) / (2 * a)
		fmt.Printf("Roots of equation are %f and %f", r1, r2)
	} else if d == 0 {
		r1 := -b / (2 * a)
		r2 := -b / (2 * a)
		fmt.Printf("Roots of equation are %g and %g", r1, r2)
	} else if d < 0 {
		r1 := -b / (2 * a)
		r2 := -b / (2 * a)
		imaginary := math.Sqrt(-d) / 2 * a
		fmt.Println("Roots of equation are root1 = ", r1, "+", imaginary, "and root2", r2, "-", imaginary)
	}

}
