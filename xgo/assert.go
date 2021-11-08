package main

import "fmt"

func main() {
	var c interface{}
	c = nil

	t, ok := c.(int64)
	if !ok {
		fmt.Println("断言失败")
		return
	}
	fmt.Println(t)
	return
}
