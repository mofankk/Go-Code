package main

import "fmt"

func main() {
	
	type Hello struct {
		a int
		b int
		c []int
	}
	m := make(map[int]*Hello)
	
	m[1] = &Hello{}
	
	m[1].a = 1
	m[1].b = 2	

	m[1].c = append(m[1].c, 1)
	m[1].c = append(m[1].c, 2)
	m[1].c = append(m[1].c, 3)
	fmt.Println(m[1])
}
