package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	b := 100
	var a *int

	a = &b

	fmt.Println(a, *a)

	p := Person{name: "Sean", age: 50}

	// 结构体指针的引用
	sp := p

	sp.age = 100
	fmt.Println(sp.age)
	fmt.Println(p.age) // 50

	sp2 := &p
	fmt.Println(sp2)
	fmt.Println(*sp2)

	fmt.Println(sp2.age) // 50
	sp2.age = 200

	fmt.Println(p.age) // 200

}
