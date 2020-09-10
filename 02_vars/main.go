package main

import "fmt"

// Cannot use shorthand outside of main function
var firstName string = "Jaime"

func main() {

	// Type not "required" as go will infer the type in which it is
	var age = 34
	var isCool = false
	// If isCool was a const we could not change the variable
	isCool = true

	// Shorthand version
	lastName := "Rodriguez"

	// Same types can be combined
	wifeFirstName, wifeLastName := "Krystal", "Rodriguez"

	// Floats are default float64. If float32 is wanted, it must be declared
	size := 1.3

	fmt.Println(firstName, lastName, age, isCool)
	fmt.Println(wifeFirstName, wifeLastName)
	//%T gives type (google "go fmt" for more)
	fmt.Printf("%T\n", size)
}
