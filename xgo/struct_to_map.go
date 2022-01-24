package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {

	var str string = `
{
    "a1":"hello",
    "b1":"world",
    "a2": {
	"a21": 13,
	"a22": true
	},
    "b2":true,
    "a3":14.53
}
`

	fmt.Print(str)

	var m = make(map[string]interface{})

	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(m)

	fmt.Println(m["a2"])

	type A struct {
		A21 int  `json:"a21"`
		A22 bool `json:"a22"`
	}

	var a A

	b, _ := json.Marshal(m["a2"])
	err = json.Unmarshal(b, &a)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a.A21, a.A22)

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
	astr := strings.Trim(strings.Replace(fmt.Sprint(arr), " ", ",", -1), "[]")
	fmt.Println(astr)

	ss := "000100"
	tArr := []byte(ss)
	tArr[2] = ss[3]
	fmt.Println(string(tArr))
}
