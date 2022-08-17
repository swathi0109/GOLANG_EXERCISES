package main

import "fmt"

func fullname(f_name, l_name *string) {
	if f_name == nil {
		panic("Runtime error: f_name should not be nil")
	}
	if l_name == nil {
		panic("Runtime error: l_name should not be nil")
	}
	fmt.Printf("%s %s\n", *f_name, *l_name)
	fmt.Println("Returned normally from fullname")
}

func main() {
	first_name := "swathi"
	last_name := "shekar"
	fullname(&first_name, &last_name)
	fmt.Println("Returned normally from main\n")

	var a [3]int = [3]int{1, 2, 3}
	fmt.Println(a[3])
}
