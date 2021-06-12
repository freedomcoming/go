package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	// 声明链表
	l := list.New()

	// 数据添加到尾部
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)


	six := l.PushBack(6)
	l.Remove(six) // 删除6这个节点
	fmt.Println(reflect.TypeOf(l))



	// 遍历
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.Value)
	}
}
