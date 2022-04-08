package main

import "fmt"

type T1 struct {
	Name string
	Age  int
}

func main() {
	t1 := &T1{
		Name: "hello",
		Age:  12,
	}
	t1 = nil
	fmt.Printf(t1.Name)

}
