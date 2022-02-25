package main

import (
	"sync"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		// 可以编译，有并发问题，f函数可能会被执行多次
		// 在多核CPU中，因为CPU缓存会导致多个核心中变量值不同步。
		f()
	}
}
