package main

import "fmt"

func fact(n int) int {
	if n <= 0 {
		return 1
	}
	return fact(n-2) + fact(n-1)
}

func main() {
	fmt.Println(fact(10))
}
