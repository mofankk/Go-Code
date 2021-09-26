package main

import (
	"encoding/json"
	"fmt"
)

type dog struct {
	Name string `json:"name" db:"name"`
	Year int64  `json:"year"`
}

func main() {
	d1 := dog{
		Name: "布迪",
		Year: 2020,
	}

	// 序列化
	b, err := json.Marshal(d1)
	if err != nil {
		fmt.Printf("报错了%v \n", err)
		return
	}
	fmt.Printf("%#v\n", string(b))
	fmt.Println(string(b))

	// 反序列化
	var p dog
	err = json.Unmarshal([]byte(b), &p)
	if err != nil {
		fmt.Printf("报错了%v \n", err)
		return
	}
	fmt.Println(p)
	fmt.Printf("%#v \n", p)
}
