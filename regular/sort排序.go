package main

import (
	"fmt"
	"sort"
)

func main() {

	stars := []string{"c", "a", "b"}
	sort.Strings(stars)
	fmt.Println("Strings:", stars)

	its := []int{7, 2, 4}
	sort.Ints(its)
	fmt.Println("Its:   ", its)

	s := sort.IntsAreSorted(its)

	fmt.Println("Sorted: ", s)
}