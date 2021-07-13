package main

import (
	"fmt"
	"reflect"
)

func main() {

	a := 1

	v := reflect.ValueOf(a)

	fmt.Println("v Type:", v.Type())
	fmt.Println("v CanSet:", v.CanSet())
	//v.SetInt(222)
	fmt.Println("v IsValid:", v.IsValid())

	v = reflect.ValueOf(&a)
	fmt.Println("v Type:", v.Type())
	fmt.Println("v CanSet:", v.CanSet())

	v = v.Elem() // element value
	fmt.Println("v Type:", v.Type())
	fmt.Println("v CanSet:", v.CanSet())

	// set
	v.SetInt(2)
	fmt.Println("after set, v:", v)

	newValue := reflect.ValueOf(3)
	v.Set(newValue)
	fmt.Println("after set, v:", v)

}
