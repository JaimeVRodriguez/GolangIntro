package main

import (
	"fmt"
	// Bringing in additional packages
	"math"
	// Bring in a package I created
	// To import it must match the directory tree it is located in
	"github.com/JaimeVRodriguez/go_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))

	// Using our created package
	fmt.Println(strutil.Reverse("olleh"))

}
