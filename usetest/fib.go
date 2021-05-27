package main

import (
	"fmt"
	"strings"
	"time"
)

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}


func GetData() {
	tNow := time.Now()
	var d []string
	for i := 6; i >= 0; i-- {
		xNow := tNow.AddDate(0, 0, -1 * i).Format("2006-1-2")
		x := strings.Split(xNow, "-")
		v := x[1] + "月" + x[2] + "日"
		d = append(d, v)
	}
	for k, v := range d {
		fmt.Println(k, v)
	}
}

