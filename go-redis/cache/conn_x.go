package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var cli1 *redis.Client

func Conn1() {
	rdb := redis.NewClient(&redis.Options{
		Addr:         "123.57.24.8:60013",
		Password:     "Robot$Cx330.@", // no password set
		DB:           0,               // use default DB
		PoolSize:     10,
		MinIdleConns: 6,
	})
	if rdb == nil {
		fmt.Errorf("get conn err \n")
	}

	status := rdb.PoolStats()
	fmt.Printf("Total Conns: %d \n ", status.TotalConns)
	fmt.Printf("Idle  Conns: %d \n ", status.IdleConns)
	fmt.Printf("Hits       : %d \n ", status.Hits)

	cli1 = rdb

}

func GetConn1() *redis.Conn {
	ctx := context.Background()
	return cli1.Conn(ctx)
}
