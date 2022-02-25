package main

import "fmt"

const (
	aa = iota
	bb = iota
)
const (
	name = "menglu"
	cc   = iota
	dd   = iota
)

const (
	a  = iota //0
	b         //1
	c         //2
	d  = "ha" //独立值，iota += 1
	e         //"ha"   iota += 1
	f  = 100  //iota +=1
	g         //100  iota +=1
	h  = iota //7,恢复计数
	ii        //8
)

const (
	i = 1 << iota
	j = 3 << iota
	k // 3<<2
	l // 3<<3
)

func main() {
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)
	fmt.Println("---------------------")
	fmt.Println(a, b, c, d, e, f, g, h, ii)
	fmt.Println("---------------------")
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}

/*
iota，特殊常量，可以认为是一个可以被编译器修改的常量。

iota 在 const 关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
*/

/*
运行结果:
0
1
1
2
---------------------
0 1 2 ha ha 100 100 7 8
---------------------
i= 1
j= 6
k= 12
l= 24
*/
