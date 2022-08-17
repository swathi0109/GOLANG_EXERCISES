package main

import (
	"fmt"
)

type Emp struct {
	name  string
	email string
	num   int
}

type Rectangle struct {
	l int
	b int

	geometry struct {
		area      int
		perimeter int
	}
}

func main() {
	// Method 1
	emp1 := Emp{"Ram", "ram.com", 123}
	emp2 := Emp{"Seetha", "seetha.com", 456}
	fmt.Println(emp1.name, emp2.name)

	// method 2
	var emp3 = Emp{"Sam", "sam.com", 10}
	fmt.Println(emp3)

	// Method 3 using new keyword
	emp4 := new(Emp) // emp4 is apointer to the instance of Emp
	emp4.name = "Ram"
	emp4.email = "ram.com"
	emp4.num = 12
	fmt.Println(*emp4)

	// Method 4 using pointer addressing operator
	var emp5 = &Emp{"Sam", "sam.com", 10}
	fmt.Println(emp5) // emp5 is a pointer to the instance of Emp
	fmt.Println(*emp5)

	var rect Rectangle
	rect.l = 5
	rect.b = 10
	rect.geometry.area = rect.l * rect.b
	rect.geometry.perimeter = 2 * (rect.l + rect.b)
	fmt.Println(rect.geometry.area)
	fmt.Println(rect.geometry.perimeter)
}
