package main

import "fmt"

// Inside parenthesis is the variable name with type after
// Return type follows parenthesis
func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(greeting("Jaime"))

	fmt.Println(getSum(3, 4))
}
