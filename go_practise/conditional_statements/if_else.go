package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter you age: ")
	scan.Scan()
	age, _ := strconv.ParseInt(scan.Text(), 10, 64)
	if age > 18 {
		fmt.Println("This person is major")
	} else {
		fmt.Println("This person is minor")
	}
}
