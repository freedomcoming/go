package main

import "os"

func main() {
	file, err := os.OpenFile("C:\\Users\\wxw\\Desktop\\go\\读取文件\\file.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	defer file.Close()

	if err != nil {
		return
	}
	file.WriteString("asasd\n") // 写入字符

	//	写入byte
	var st = "测试"
	file.Write([]byte(st))
}
