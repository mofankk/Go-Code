package main

import (
	"fmt"
	"geekbyte.cn/go_redis_test/pkg/redis"
	"geekbyte.cn/go_redis_test/server"
)

func main() {
	rdb, err := redis.NewRedisDataCache()
	if err != nil {
		fmt.Errorf("%v\n", err)
	}
	s := server.New(rdb)
	//for i := 0; i < 10; i++ {
	//	key := "hello_" + strconv.Itoa(i)
	//	s.SetKey(key, i, 10 * time.Minute)
	//}
	s.GetKeys()
}