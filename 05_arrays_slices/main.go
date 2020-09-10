package main

import "fmt"

func main() {

	// Arrays
	var fruitArray [2]string

	// Assign array values
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	// Declare and assign
	drinkArray := [2]string{"Water", "Mt. Dew"}

	// Slice, similar to array but does not need a predfined length
	candySlice := []string{"Red Vines", "Gushers", "Gummy Bears"}

	fmt.Println(fruitArray)
	fmt.Println(fruitArray[1])

	fmt.Println(drinkArray)

	fmt.Println(candySlice)
	fmt.Println(len(candySlice))
}
