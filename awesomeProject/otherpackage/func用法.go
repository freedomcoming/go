package otherpackage

import "fmt"

func Fun() (res int) {
	b := func(data int, data2 ...int) { // 可变参数
		fmt.Println(data, data2)
	}
	b(1010, []int{1, 2, 3}...) // 输入可变参数


	// 自执行函数

	(func() {
		fmt.Println(123)
	})()

	res = 100
	return

}


// 闭包函数 闭包不是返回函数的函数，是返回函数的函数可以读取到第一此函数内部的变量

func Closer()func()  {
	return func() {
		fmt.Println("闭包函数")
	}

}