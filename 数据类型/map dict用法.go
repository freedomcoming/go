package main

import "fmt"

func main() {

	//m := map[string]int{} // 等同以下申明
	var m map[string]int
	m = make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
	// 改变值
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// 求长度
	fmt.Println("len:", len(m))

	// 删除元素
	delete(m, "k2")
	fmt.Println("map:", m)

	prs1, prs := m["k2"]
	fmt.Println(prs1) // 不存在为0
	fmt.Println("prs:", prs)


	//
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)


	for c := range m {
		fmt.Println(m[c])
	}

}
