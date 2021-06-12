package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "string"
	fmt.Println(a)
	// 是否存在
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(len("谷歌中国"))
	// 计数
	fmt.Println(strings.Count("谷歌中国", ""))

	// 划分
	fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))
	fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))

	//
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "C"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
	fmt.Println(strings.HasSuffix("Amigo", "go"))
	fmt.Println(strings.HasSuffix("Amigo", "Ami"))
	fmt.Println(strings.HasSuffix("Amigo", ""))
}
