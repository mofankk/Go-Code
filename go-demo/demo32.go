package main

import (
	"fmt"
	"time"
)

// 32. 写出以下逻辑，要求每秒钟调用一次 proc 并保证程序不退出?
// 题目主要考察了两个知识点:
// 1. 定时执行执行任务
// 2. 捕获 panic 错误
func main() {
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()
			}
		}
	}()
	select {}
}

func proc() {
	panic("ok")
}
