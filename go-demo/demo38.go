package main

import (
	"fmt"
)

type Student struct {
	Name string
}

func main() {
	// 指针类型比较的是指针地址，非指针类型比较的是每个属性的值。
	fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"}) // true
	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})   // false
}
