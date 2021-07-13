package main

import (
	"encoding/json"
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

func MapJson() {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["Ab"] = "hello"
	m["Ac"] = 1
	m["Ad"] = 2

	b, _ := json.Marshal(m)

	type Tt struct {
		Ab 	string 	`json:"ab"`
		Ac  int  	`json:"ac"`
		Ad 	int 	`json:"ad"`
	}
	var t Tt
	e := json.Unmarshal(b, &t)
	if e != nil {
		fmt.Println(e)
	}

	fmt.Printf("%v", t)
}

