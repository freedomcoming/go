package main

import (
	"fmt"
	"reflect"
)

func reflectValue2(x interface{}) {
	v := reflect.ValueOf(x)

	fmt.Println(v.Kind())        // ptr
	fmt.Println(v.Elem().Kind()) // int

	// 传入值的地址
	v.Elem().SetInt(123)
}

func main() {
	var a int = 64
	reflectValue2(&a)
}
