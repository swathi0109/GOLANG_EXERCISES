package main

import "fmt"

func calc_notes(amount int)
func main() {
	num := 100
	var amount int
	fmt.Println("enter the amount: ")
	fmt.Scanln(&amount)
	notes := []int{500, 200, 100, 50, 20, 10, 2, 1}
	fmt.Printf("No of 10 ruppees note = %d\n", num/10)
	fmt.Printf("No of 20 ruppees note = %d\n", num/20)
	fmt.Printf("No of 50 ruppees note = %d\n", num/50)
	for i := 0; i < len(notes); i++ {
		fmt.Printf("Number of %d notes is %d\n", notes[i], amount/notes[i])
	}
}
