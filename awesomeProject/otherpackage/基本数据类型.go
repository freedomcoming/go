package otherpackage

import (
	"fmt"
	"strconv"
)

var B = 123

func IntRange() {
	var num1 uint16
	var num2 int16
	num1 = 999
	num2 = -999
	fmt.Println(num1, num2)

	var float1 float64
	float1 = 1.11111

	fmt.Println(float1)

	//	 布尔型
	var boolType bool = true
	fmt.Println(boolType)

	fmt.Printf("%T", boolType)

	fmt.Printf("%T", "123")

	//	类型转换

	string1 := "123"
	int, _ := strconv.Atoi(string1)
	fmt.Println("int是", int)

}
