package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}
type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func main() {
	// 在 golang 语言中， Student 和 *Student 是两种类型，第一个是表示 Student 本身，第二个是指向 Student 的指针。
	// var peo People = Student{}
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
