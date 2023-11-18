package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p Person) GetName() {
	fmt.Println("getName func")
}

func reflectStruct(x interface{}) {

	// 类型获取
	t := reflect.TypeOf(x)

	fmt.Println(t.Field(0).Name, t.Field(0).Type, t.Field(0).Tag.Get("json"))
	fmt.Println(t.FieldByName("Age"))
	// 值获取
	v := reflect.ValueOf(x)
	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.Field(1).Int())

	//	 方法获取
	// 多个方法 顺序按照 ASCII码有关系
	// 同样支持MethodByName
	mm := t.Method(0)
	fmt.Println(mm.Name) // GetName

	//	 执行方法
	v.MethodByName("GetName").Call(nil)
}
func main() {
	a := Person{Name: "karl", Age: 12}

	reflectStruct(a)

	reflectStructModifyValue(&a)

	fmt.Println(a)
}

func reflectStructModifyValue(x interface{}) {
	v := reflect.ValueOf(x)
	name := v.Elem().FieldByName("Name")
	name.SetString("123")

}
