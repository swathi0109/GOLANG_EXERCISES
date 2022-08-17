package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the year you were born: ")
	scanner.Scan()
	input2, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d years at the end of 2022", 2022-input2)
}
