package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("C:\\Users\\wxw\\Desktop\\go\\读取文件\\fileRead.go")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	if err != nil {
		fmt.Println(err)
	}
	// 切片
	var strSlice []byte
	var temp = make([]byte, 128)

	// 循环读取
	for {
		n, err2 := file.Read(temp)

		if err2 == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		strSlice = append(strSlice, temp[:n]...)

		if err2 != nil {
			return
		}

	}
	fmt.Println(string(strSlice))

}
