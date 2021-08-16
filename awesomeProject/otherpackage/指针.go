package otherpackage

import "fmt"

func Zhizhen() {
	a := 123

	fmt.Println(a)

	var b *int
	b = &a
	fmt.Println(a, b)
	*b = 999
	fmt.Println(a, b, *b, &a, &b)

	var name = "123"
	nameP := &name // 类型不能改变 快速建立指针
	*nameP = "fuck 123"
	fmt.Println(name)
}
