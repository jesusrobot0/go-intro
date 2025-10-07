package main

import "fmt"

func main() {

	// For
	age := 0
	// for life := 0; life < 120; life++ {
	// 	age++
	// 	fmt.Println("Este año tienes ", age, " años")
	// }

	// While
	// age = 0
	// for age < 120 {
	// 	age++
	// }
	// fmt.Println("Este año tienes ", age, " años")
	
	// Infinite loop
	for {
		age++
		
		if age == 120 {
			fmt.Println("¡Ya eres un anciano!")
			break
		}
	}
}