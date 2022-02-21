package main

import "fmt"

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		ch <- vs[i](name)
	}

	for i := range vs {
		go fn(i)
	}

	return <-ch
}

// 16. 请说出下面的代码存在什么问题?
func main() {
	str := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(str)
}

/*
解释:
TODO 为什么为很可能打印 111func4 ?
依据4个goroutine的启动后执行效率，很可能打印111func4，但其他的111func*也可能先执行，exec只会返回一 条信息。
*/
