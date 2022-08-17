package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the text: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("Your typed: ", input)
}
