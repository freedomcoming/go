package main

import (
	"fmt"
	"goProject/regular/packagenew"
)

func main() {
	// 包的使用 目录中才能建立包  可以别名 cool .  大写变量才能引用 小写是私有的
	//fmt.Println(cool.A)
	//fmt.Println(A)
	packagenew.Test1()
	var a = "initial"
	aa := "aa"
	fmt.Println(a, aa)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	ff := "apple2"
	fmt.Println(ff)

	var fff string = "apple3"
	fmt.Println(fff)
}
