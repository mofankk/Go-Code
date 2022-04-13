package main

import (
	"fmt"
)

type Student struct {
	Name string
}

type Tong struct {
	Age int
}

func main() {
	ps := &Student{}
	s := Student{}
	t := Tong{}
	whot(ps)
	whot(s)
	whot(t)
}

func whot(ty interface{}) {
	fmt.Println("hello")

	switch ty.(type) {
	default:
		fmt.Println("type not support")
	case *Student:
		fmt.Println("Student point")
	case Student:
		fmt.Println("Student")
	case Tong:
		fmt.Println("大桐子")
	}

}
