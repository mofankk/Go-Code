package main

import "fmt"

type People struct{}

func (p *People) String() string {
	// Err: fmt.Sprintf format %v with arg p causes recursive
	return fmt.Sprintf("%v", p)
}

func main() {
	p := &People{}
	fmt.Println(p.String())
}

/*
解析：
在golang中String()string方法实际上是实现了String的接口的，该接口定义在fmt/print.go 中。
在使用 fmt 包中的打印方法时，如果类型实现了这个接口，会直接调用。而题目中打印 p 的时候会直接调用 p 实现的 String() 方法，然后就产生了循环调用。
*/
