package main

import (
	"runtime"
	"strconv"
)

func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 5; i++ {

		go say("Hello World: " + strconv.Itoa(i))
	}


}

func say(s string) {
	println(s)
}
