package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type People struct {
	// Warn: struct field name has json tag but is not exported
	name string `json:"name"`
}

// 9. 写出打印的结果。
// 打印结果为：people {}
func main() {
	js := `{
		"name": "ll"
		}`
	var people People
	err := json.Unmarshal([]byte(js), &people)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("people", people)
}
