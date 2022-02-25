package main

import (
	"fmt"
)

func main() {
	str1 := []string{"a", "b", "c"}
	// slice底层是数组,这样 str2 与 str1 共用一个数组，修改str2会影响到str1
	str2 := str1[1:]
	str2[1] = "new"
	fmt.Println(str1)
	// 数组append发生扩容，此时str2是一个新的数组
	//str2 = append(str2, "z", "x", "y")
	str2 = append(str2, "z")
	fmt.Println(str1)

	var str3 = make([]string, 3)
	// copy(dest, src) Go语言的内置函数 copy() 可以将一个数组切片复制到另一个数组切片中，如果加入的两个数组切片不一样大，就会按照其中较小的那个数组切片的元素个数进行复制。
	copy(str3, str1)
	str3[1] = "hello"
	fmt.Println(str1)
	fmt.Println(str3)
}

/*
运行结果:
[a b new]
[a b new]
[a b new]
[a hello new]
*/
