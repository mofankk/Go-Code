package main

import "fmt"

func main() {
	type ss struct {
		age int
	}
	m := make(map[string]*ss)
	m["t"] = &ss{age: 11}
	fmt.Println(m["t"].age)

	age := m["t"]
	age.age = 20

	fmt.Println(m["t"].age)

	m1 := make(map[string]ss)
	m1["t"] = ss{age: 11}
	fmt.Println(m1["t"].age)

	age1 := m1["t"]
	age1.age = 20

	fmt.Println(m1["t"].age)

}

// map 返回指针到一个变量的时候，修改这个变量会影响到map里的值

/*
执行结果
11
20
11
11
*/
