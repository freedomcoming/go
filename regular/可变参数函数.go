package main

import "fmt"

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}

	fmt.Println()

	sum(nums...)
	//fmt.Println(nums)
	nums2 := []int{2, 3, 4, 5}
	sum(nums2...)
}
