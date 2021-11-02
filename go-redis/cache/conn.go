package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var cli *redis.Client

func Conn() {
	rdb := redis.NewClient(&redis.Options{
		Addr:	  "192.168.4.112:6379",
		Password: "Hello", // no password set
		DB:		  0,  // use default DB
		PoolSize: 10,
		MinIdleConns: 6,
	})
	if rdb == nil {
		fmt.Errorf("get conn err \n")
	}

	status := rdb.PoolStats()
	fmt.Printf("Total Conns: %d \n ", status.TotalConns)
	fmt.Printf("Idle  Conns: %d \n ", status.IdleConns)
	fmt.Printf("Hits       : %d \n ", status.Hits)

	cli = rdb

}

func GetConn() *redis.Conn {
	ctx := context.Background()
	return cli.Conn(ctx)
}
