package main

import "fmt"

type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"Mofan"}}

	m1 := map[string]*Student{"kk": {"World"}}
	// 这一行会出错
	// map的value本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在map改变时会变得无效。 故如果需要修改map值，可以将 map 中的非指针类型 value ，修改为指针类型，比如使用 map[string]*Student .
	//m["people"].name = "hello"
	fmt.Println(m1["kk"])
	fmt.Println(m1["kk"].name)
	m1["kk"].name = "hello, world"
	//m1["kk"] = &Student{"fine"}
	fmt.Println(m1["kk"])

	fmt.Println(m["people"].name)
}
