package main

import "fmt"

func fact(n int16) int16 {
	if (n == 0) || (n == 1) {
		return 1
	}
	return (n * fact(n-1))
}

func main() {
	var n int16
	fmt.Println("Enter the number to find factorial")
	fmt.Scanln(&n)
	factorial := fact(n)
	fmt.Printf("factorial of %d is %d", n, factorial)
}
