package main

import (
	"fmt"
	"reflect"
)

// 定义自己的类型

type mint int

// 空接口的代表任意类型
func reflectFn(x interface{}) {

	v := reflect.TypeOf(x)
	fmt.Println(v)

	fmt.Printf("name is %v,kind is %v \n", v.Name(), v.Kind())

}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)

	// v的类型是reflect.value
	fmt.Println(v)
	//	 v.Int()

	fmt.Println(v.Kind()) // int

	switch v.Kind() {

	case reflect.Int:
		fmt.Println(v.Int())

	default:
		fmt.Println("other")

	}

	fmt.Println(v.Int() + 1)
}

func main() {

	a := 100
	reflectFn(a)

	reflectValue(a)

	//aa := struct {
	//}{}
	//
	//reflectFn(aa)
	//
	//var my mint
	//
	//reflectFn(my)
	//
	//reflectFn(&a) // *int 指针类型
	//
	////	数组和slice
	//i := [3]int{1, 2, 3} // array
	//j := []int{3, 3, 3}  // slice
	//
	//reflectFn(i)
	//reflectFn(j)
}
