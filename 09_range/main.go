package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	// i is index
	// id is the iterator
	// ids is array/slice name
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	// Use filler _
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids
	total := 0
	for _, id := range ids {
		total += id
	}
	fmt.Println("Total:", total)

	// Range w/ map
	employeeID := map[string]string{"1": "Bob", "2": "Kyle", "3": "Mike"}

	// k = key, v = value (can use any letters)
	for k, v := range employeeID {
		fmt.Printf("%s: %s\n", k, v)
	}

	for _, v := range employeeID {
		fmt.Printf("Name: %s\n", v)
	}

}
