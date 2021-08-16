package otherpackage

import "fmt"

func Map() (res1 int, res2 string) { // 什么输出返回值
	m1 := map[string]string{}
	m2 := make(map[string]interface{})
	m1["name"] = "wxw"
	m1["sex"] = "男"
	fmt.Println("------")
	fmt.Println(m1["name"], m2)

	//	value 不限制类型 空接口 interface

	m2["name"] = 100
	m2["sex"] = "女"
	m2["list"] = []int{1, 2, 3}

	fmt.Println(m1, m2)

	//	delete

	delete(m2, "sex")
	fmt.Println(m1, m2)

	//	循环输出

	for k, v := range m2 {
		fmt.Println(k, v)
	}
	//return 0,"ok"
	// res1 和 res2使用
	res1 = 1
	res2 = "ok"
	return
}
