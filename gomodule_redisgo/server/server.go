package server

import (
	"fmt"
	"geekbyte.cn/go_redis_test/pkg/redis"
	"time"
)

type Server interface {
	SetKey(key string, val interface{}, timeout time.Duration)
	GetKeys()
}

type service struct {
	rdb *redis.Redis
}

func New(rdb *redis.Redis) Server {
	return &service{
		rdb,
	}
}

func (s *service) SetKey(key string, val interface{}, timeout time.Duration) {

	err := s.rdb.Set(key, val, timeout)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *service) GetKeys() {
	keys, err := s.rdb.GetKeys("hell*")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range keys {
		fmt.Println(v)
	}
}