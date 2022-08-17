//Recover function will stop the panicking sequence execution by restoring the normal execution and retrieves the error message
//when called inside the defer function

package main

import "fmt"

func recoverfullname() {
	r := recover()
	if r != nil {
		fmt.Printf("Recovered from ", r)
	}

}

func fullname(f_name, l_name *string) {
	defer recoverfullname()

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
	//last_name := "shekar"
	fullname(&first_name, nil)
	fmt.Printf("Returned normally from main\n")
}
