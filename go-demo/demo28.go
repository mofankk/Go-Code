package main

import (
	"fmt"
)

type People interface {
	Show()
}
type Student struct{}

func (stu *Student) Show() {
}
func live() People {
	var stu *Student
	return stu
}
func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

/*
运行结果:
BBBBBBB

*Student 的定义后本身没有初始化值，所以 *Student 是 nil 的，但是 *Student 实 现了 People 接口，接口不为 nil 。
*/
