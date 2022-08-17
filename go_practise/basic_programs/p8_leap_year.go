package main

import "fmt"

func main() {
	var year int
	fmt.Println("Enter the year")
	fmt.Scanln(&year)
	if (year%4 == 0) && (year%100 != 0) || (year%400 == 0) {
		fmt.Printf("%d is a leap year", year)
	} else {
		fmt.Printf("%d is not a leap year", year)
	}
}
