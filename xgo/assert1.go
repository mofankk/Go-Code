package main

import "fmt"

func main() {
	var c interface{}
	c = 1

	// 竟然还能这么玩
	switch t := c.(type) {
	case int:
		fmt.Println("int", t)
		t = 2
		fmt.Println("t == ", t)
	}

}
