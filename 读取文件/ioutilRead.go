package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	byteStr, err := ioutil.ReadFile("C:\\Users\\wxw\\Desktop\\go\\读取文件\\fileRead.go")
	if err != nil {
		return
	}
	fmt.Println(string(byteStr))
}
