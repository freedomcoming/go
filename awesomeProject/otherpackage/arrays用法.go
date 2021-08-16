package otherpackage

import "fmt"

func Arrays() {

	//var a [5]int
	a := [5]int{}
	fmt.Println("emp:", a)

	//var bb [3]string
	bb := make([]int, 3)
	fmt.Println(bb)

	// 另一种写法

	cc := new([3]int)
	fmt.Println("cc-", cc)

	// 输出序列

	strings := []string{"a", "b", "c"}

	for i, v := range strings {
		fmt.Println(i, "-", v)
	}

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 多维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// 不定长数组

	LongArray := [...]int{123, 456}

	fmt.Println(len(LongArray))
	fmt.Println("cap", cap(LongArray)) // 容积

	//	copy
	//  被copy和被赋值的必须是slice
	// 切片不是数组 不固定长度的数组
	newArray1 := [5]int{1, 2, 3, 4, 5}
	var newArray2 []int
	copy(newArray2, newArray1[1:])
	fmt.Println(newArray1, newArray2)

	// 切片使用
	a[4] = 100
	fmt.Println("set:", a)
	newArray := a[1:]
	newArray = append(newArray, 5) // 切片后可以append
	newArray = append(newArray, 5)
	newArray = append(newArray, 5)
	newArray = append(newArray, 5)
	newArray = append(newArray, 5)
	// 容积达到临界会翻倍扩容
	fmt.Println(len(newArray), cap(newArray))
	fmt.Println("get:", newArray)

	//	 切片创建
	//var sliceArray []int
	//or
	sliceArray := make([]int,4)
	fmt.Println(sliceArray)
}
