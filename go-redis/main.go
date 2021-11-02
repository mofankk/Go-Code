package main

import (
	"fmt"
	"context"
	"geekbyte.cn/go-redis/cache"
)

func main() {
	cache.Conn()

	cli := cache.GetConn()
	cli.Set(context.Background(), "a", "c", 0)
	t, err := cli.Get(context.Background(), "a").Result()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(t)
	cli.Close()

	cli1 := cache.GetConn()
	cli1.Set(context.Background(), "b", "c", 0)
	t, err = cli1.Get(context.Background(), "b").Result()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(t)
	cli1.Close()

	cli2 := cache.GetConn()
	cli2.Set(context.Background(), "c", "c", 0)
	t, err = cli2.Get(context.Background(), "c").Result()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(t)
	cli2.Close()
}


func runA() {
}

