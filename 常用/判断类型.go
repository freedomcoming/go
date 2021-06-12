package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int64 = 20
	fmt.Println("type:", reflect.TypeOf(x))
}
