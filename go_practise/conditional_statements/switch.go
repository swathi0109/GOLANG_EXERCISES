package main

import "fmt"

func main() {
	flavour := "pista"
	switch flavour {
	case "vanilla":
		fmt.Println("Vanilla")
	case "mango":
		fmt.Println("Mango")
	default:
		fmt.Println("Flavour not specified")
	}
}
