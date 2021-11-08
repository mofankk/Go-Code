package main

import (
	"context"

	"geekbyte.cn/go-redis/cache"
)

func main() {
	cache.Conn1()

	cli := cache.GetConn1()
	m := make(map[string]string)

	m["cui1"] = "cui1"
	m["cui2"] = "cui2"

	cli.HSet(context.Background(), "sn110011", m)

	//t, err := cli.Get(context.Background(), "a").Result()
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//}
	//fmt.Println(t)
	//cli.Close()

	//cli1 := cache.GetConn()
	//cli1.Set(context.Background(), "b", "c", 0)
	//t, err = cli1.Get(context.Background(), "b").Result()
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//}
	//fmt.Println(t)
	//cli1.Close()

	//cli2 := cache.GetConn()
	//cli2.Set(context.Background(), "c", "c", 0)
	//t, err = cli2.Get(context.Background(), "c").Result()
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//}
	//fmt.Println(t)
	//cli2.Close()
}
