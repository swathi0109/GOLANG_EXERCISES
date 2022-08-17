//map is represented as key:value pairs
package main

import "fmt"

func main() {
	//map creation method 1
	var capitalmap map[string]string
	capitalmap = make(map[string]string)
	capitalmap["India"] = "New Delhi"
	capitalmap["France"] = "Paris"
	capitalmap["Swiss"] = "Bern"

	for country := range capitalmap {
		fmt.Println("Capital of", country, "is", capitalmap[country])
	}
	fmt.Println()

	capital, ok := capitalmap["Agra"]
	if ok {
		fmt.Println("Captial of India is ", capital)
	} else {
		fmt.Println("Country not found")
	}
	fmt.Println()

	//map delete
	delete(capitalmap, "France")
	fmt.Println(capitalmap)
	fmt.Println()

	//map creation method 2
	map1 := map[int]string{
		1: "a\n",
		2: "b\n",
		3: "abc",
	}
	fmt.Println(map1)
	fmt.Println()

	//map creation method 3 using make
	map2 := make(map[string]int) // string-datatype of key, int-datatype of value
	map2["Ram"] = 1
	map2["Raj"] = 2
	map2["Rim"] = 3
	fmt.Println(map2)
	fmt.Println()

	//iteration of map using for loop
	for id, elem := range map2 {
		fmt.Println(id, elem)
	}
	fmt.Println()

	//addition of new key-vaue pairs
	map2["Sam"] = 4
	fmt.Println(map2)

}
