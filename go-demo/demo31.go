package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Ban struct {
	visitIPs map[string]time.Time
	lock     sync.Mutex
}

func NewBan(ctx context.Context) *Ban {
	ban := &Ban{visitIPs: make(map[string]time.Time)}
	go func() {
		timer := time.NewTimer(time.Minute * 1)
		select {
		// 每分钟清理过期IP
		case <-timer.C:
			ban.lock.Lock()
			for k, v := range ban.visitIPs {
				if time.Now().Sub(v) > time.Minute*1 {
					delete(ban.visitIPs, k)
				}
				ban.lock.Unlock()
				timer.Reset(time.Minute * 1)
			}
		case <-ctx.Done():
			return
		}
	}()
	return ban
}

func (o *Ban) visit(ip string) bool {
	o.lock.Lock()
	defer o.lock.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}

func main() {
	success := int64(0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(100 * 1000)

	ban := NewBan(ctx)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(t int) {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.10.%d", t)
				if ban.visit(ip) {
					// 多核CPU修改int值在极端情况下会有不同步情况，因此需要原子性的修改int值
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}
	}
	wg.Wait()
	fmt.Printf("success: %d", success)
}
