package main

import "fmt"

func main() {
	type void struct{}
	var member void

	set := make(map[string]void) // New empty set
	set["Foo"] = member          // Add
	for k := range set {         // Loop
		fmt.Println(k)
	}
	delete(set, "Foo") // Delete
	size := len(set)   // Size
	fmt.Println(size)
	_, exists := set["Foo"] // Membership
	fmt.Println(exists)
}
