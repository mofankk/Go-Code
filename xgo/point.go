package main

import "fmt"

type S1 struct {
	Name string
	S2   *S2
}

type S2 struct {
	Age  int
	Addr string
}

func main() {
	s1 := &S1{}
	//s1.S2 = new(S2)
	s1.Name = "dd"
	//s1.S2.Age = 2
	//s1.S2.Addr = "cc"

	s1.S2 = &S2{
		Age:  1,
		Addr: "cc",
	}

	fmt.Println(s1.S2.Addr)

}
