package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now().Unix()
	for num := 2; num < 100; num++ {
		flag := false
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = true
				break
			}
		}
		if flag {
			fmt.Printf("%d 是素数 \n", num)
		}

	}
	end := time.Now().Unix()

	fmt.Println(end - start)
}
