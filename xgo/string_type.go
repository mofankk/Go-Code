package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "hello, world"
	for i, v := range str {
		fmt.Println(reflect.TypeOf(str[i]).Name())
		fmt.Println(reflect.TypeOf(v).Name())
		break
	}
}
