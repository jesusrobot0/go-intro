package main

import "fmt"

func main() {
	fruitsCodes := make(map[string]int)

	fruitsCodes["apple"] = 1
	fruitsCodes["banana"] = 2
	fruitsCodes["cherry"] = 3

	fmt.Println(fruitsCodes)
	fmt.Println(fruitsCodes["apple"])

	delete(fruitsCodes, "banana")
	fmt.Println( fruitsCodes)
	
	_, exists := fruitsCodes["banana"]

	fmt.Println("exists banana: ", exists)

	frutas := map[string]int{
		"apple": 1,
		"banana": 2,
		"cherry": 3,
	}

	fmt.Println(frutas)
}
