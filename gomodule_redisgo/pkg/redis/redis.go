package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"geekbyte.cn/go_redis_test/config"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	DB0 = 0
)

type Redis struct {
	conn *redis.Pool
}

func NewRedisDataCache() (*Redis, error) {
	cfg, err := config.New("redis")
	if err != nil {
		return nil, err
	}
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.GetString("redis.host"), redis.DialDatabase(cfg.GetInt("redis.db")), redis.DialPassword(cfg.GetString("password")))
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err = c.Do("PING")
			return err
		},
		MaxIdle:     cfg.GetInt("redis.maxIdle"),
		MaxActive:   cfg.GetInt("redis.maxActive"),
		IdleTimeout: time.Second * time.Duration(cfg.GetInt64("redis.idleTimeout")),
	}
	return &Redis{pool}, nil
}

func (r *Redis) SetRedisPoll(pool *redis.Pool) {
	r.conn = pool
}

func (r *Redis) SetConn(conn *redis.Pool) {
	r.conn = conn
}

func (r *Redis) DeleteWithPattern(pattern string) error {
	conn := r.conn.Get()
	defer conn.Close()

	keys, err := r.GetKeys(pattern)
	if err != nil {
		return err
	}
	for _, v := range keys {
		conn.Send("DEL", v)
	}
	conn.Flush()
	_, err = conn.Receive()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) GetKeys(pattern string) ([]string, error) {
	conn := r.conn.Get()
	defer conn.Close()

	iter := 0
	var keys []string
	for {
		arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
		if err != nil {
			return keys, fmt.Errorf("error retrieving %s keys", pattern)
		}
		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)
		if iter == 0 {
			break
		}
	}
	return keys, nil
}

func (r *Redis) Set(key string, val interface{}, timeout time.Duration) (err error) {
	conn := r.conn.Get()
	defer conn.Close()

	var data []byte
	if data, err = json.Marshal(val); err != nil {
		return
	}
	_, err = conn.Do("SETEX", key, int64(timeout/time.Second), data)
	return
}

func (r *Redis) Get(key string) (interface{}, error) {
	conn := r.conn.Get()
	defer conn.Close()

	var (
		data []byte
		err error
	)
	data, err = redis.Bytes(conn.Do("GET", key))
	if errors.Is(err, redis.ErrNil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var reply interface{}
	if err = json.Unmarshal(data, &reply); err != nil {
		return nil, err
	}
	return reply, nil
}