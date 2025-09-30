package main

import "fmt"

func main() {

	// Arrays
	fruits := [5]string{"Apple", "Orange", "Grape", "Banana"}

	
	fmt.Printf("Mis frutas favoritas son: %d \n", len(fruits))
	fmt.Printf("Mis frutas favoritas son: %s \n", fruits[4])
	fruits[4] = "Watermelon" 
	fmt.Printf("Mis frutas favoritas son: %s \n", fruits[4])

	// Slice

	ages := []int {22, 25, 62}

	fmt.Printf("Size: %d, value: %v \n", len(ages), ages)

	ages = append(ages, 3, 5)

	fmt.Printf("Size: %d, value: %v \n", len(ages), ages)

	humanAges := ages[:3]
	petAges := ages[3:]
	fmt.Printf("Humans: Size: %d, value: %v \n", len(humanAges), humanAges)
	fmt.Printf("Petsa: Size: %d, value: %v \n", len(petAges), petAges)


	{
		ages := []int{22, 25, 62}
		humanAges := ages[:2]

		humanAges = append(humanAges, 999)

		fmt.Println("humanAges:", humanAges)
		fmt.Println("ages:", ages)
	}

	// sobre el GC

	{
		var results []int
		// "No importa, el GC lo limpia"
		for i := 0; i < 1000000; i++ {
				results = append(results, i)
		}
	}

	{
		// ✅ Creas menos trabajo para el GC
		results := make([]int, 0, 1000000)
		for i := 0; i < 1000000; i++ {
				results = append(results, i)  // Un solo array
		}
	}

}