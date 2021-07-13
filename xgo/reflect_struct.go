package main

import (
	"fmt"
	"reflect"
)

type Orange struct {
	Size int
}

func main() {

	a := Orange{99}

	v := reflect.ValueOf(a)
	fmt.Println("v:", v)

	fmt.Println("v Type:", v.Type())
	fmt.Println("v CanSet:", v.CanSet())

	v = reflect.ValueOf(&a)
	fmt.Println("v:", v)

	fmt.Println("v Type:", v.Type())
	fmt.Println("v CanSet:", v.CanSet())

	//element
	v = v.Elem()

	size := v.FieldByName("Size")
	fmt.Println("size CanSet:", size.CanSet())

	size.SetInt(88)

	fmt.Println("after set:", v)

}
