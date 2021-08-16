package otherpackage

import "fmt"

func InterfaceUse() {
	// 泛型参数，
	(func(a interface{}) {
		fmt.Println(a)
	})([3]int{1, 2, 3})
}
