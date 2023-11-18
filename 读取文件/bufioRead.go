package main

import (
	"bufio"
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
	reader := bufio.NewReader(file)

	var fileStr string

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			fileStr += line
			break
		}

		fileStr += line

	}

	fmt.Println(fileStr)

}
