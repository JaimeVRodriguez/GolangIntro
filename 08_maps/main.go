package main

import "fmt"

func main() {
	// Define map
	// First type is key, second type is value
	emails := make(map[string]string)

	// Assign key/values
	emails["Bob"] = "bob@gmail.com"
	emails["Kyle"] = "kyle@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(len(emails))
	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	// Delete
	// First value is map name, second value key to be deleted
	delete(emails, "Bob")
	fmt.Println(emails)

	//****************************************

	// Declare and assign

	employeeID := map[string]string{"1": "Bob", "2": "Kyle"}

	fmt.Println(employeeID)
}
