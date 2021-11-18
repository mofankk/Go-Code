package main

import "fmt"

func main() {
	fmt.Println("main")
	defer func() {
		fmt.Println("main_end")
	}()

	hello()
}

func hello() {
	fmt.Println("hello_1")
	defer func() {
		// 如果没有捕获错误的地方程序会崩溃
		//if err := recover(); err != nil {
		//	fmt.Println("hello_1, rerecrecover")
		//}
		fmt.Println("world_1")
	}()

	hello2()
}

func hello2() {
	fmt.Println("hello_2")
	defer func() {
		fmt.Println("world_2")
	}()

	a, b := 666, 0
	c := a / b
	fmt.Println(c)
}
