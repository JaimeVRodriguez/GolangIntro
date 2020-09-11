package main

import (
	"fmt"
)

func main() {

	x := 5
	y := 10

	// Parenthesis are not required
	// Dependent on organization practices
	if x < y {
		// Printf is for formatting
		// %d is a "placeholder" for variables to be inserted
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//****************************************

	color := "red"

	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is not red or blue")
	}

	//****************************************

	// Switch statement
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not red or blue")
	}

}
