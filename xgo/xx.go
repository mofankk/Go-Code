package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Abc struct {
	Name string
	Age  int
	Addr string
	Ok   *bool
}

func main() {

	abc := make([]Abc, 5)

	for i := 0; i < 5; i++ {
		abc[i].Name = strconv.Itoa(i) + "-cc"
		abc[i].Age = i
		abc[i].Addr = strconv.Itoa(i)

	}

	for i := range abc {
		rt := reflect.ValueOf(&abc[i]).Elem()
		if abc[i].Name == "1-cc" {
			//rt.FieldByName("Ok").SetBool(true)
			c1 := rt.FieldByName("Ok")
			if c1.IsValid() {
				fmt.Println("isvalid")
				if c1.CanSet() {
					fmt.Println("can set")
				}
				ok := true
				cx := reflect.ValueOf(&ok)
				c1.Set(cx)
				//cc := c1.Elem()
				//cc.SetBool(true)
				//c1.SetBool(true)
			}
		} else {
			var x int64 = 100

			c2 := rt.FieldByName("Agee")
			if c2.IsValid() {
				fmt.Println("c2 valid")
				c2.SetInt(x)
			}
			c3 := rt.FieldByName("Addrr")
			if c3.IsValid() {
				fmt.Println("c3 valid")
				c3.SetString("Hello,aaaaa")
			}
		}
	}

	for i := range abc {
		fmt.Println(abc[i])
	}
}
