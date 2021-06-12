package main

import (
	"fmt"
	"github.com/asmcos/requests"
)

func main() {

	req := requests.Requests()
	resp, _ := req.Get("http://httpbin.org/ip")
	//println(resp.Text())
	fmt.Println(resp.Text())
	fmt.Println(resp.R.StatusCode)
	fmt.Println(resp.R.Header["Content-Type"])
}
