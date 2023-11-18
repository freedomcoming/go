package main

import "fmt"

func main() {

	onlyRead := make(<-chan int, 2)

	onlyWrite := make(chan<- int, 2)

	fmt.Println(onlyWrite, onlyRead)

}
