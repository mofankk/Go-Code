package main

import (
	"docker/api"
	"fmt"
	"net/http"
)

func main() {

	h := api.Hello{}

	http.HandleFunc("/hello", h.SayHello)

	fmt.Println("服务启动，端口为:8001")
	_ = http.ListenAndServe(":8001", nil)

}
