package main

import (
	"fmt"
)

func main() {
	// 数组只能与相同维度、⻓度以及类型的其他数组比较，切片之间不能直接比较。
	fmt.Println([...]string{"1"} == [...]string{"1"}) // true
	//fmt.Println([]string{"1"} == []string{"1"}) // error
}
