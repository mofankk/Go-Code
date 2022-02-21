package main

import (
	"fmt"
	"sync/atomic"
)

var value int32

// 12. 请说明下面代码书写是否正确。
func main() {
	value = 2
	setValue(1)
}

func setValue(delta int32) {
	// 调用atomic.CompareAndSwapInt32()无需使用循环
	// 该方法的作用是取到给定内存地址下的值，如果和给定的旧值相等，则用新值替换旧值，并返回true，如果和给定旧值不想等，则返回false
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, v+delta) {
			break
		}
	}
	fmt.Println(atomic.LoadInt32(&value))
}
