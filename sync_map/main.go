package main

import (
	"fmt"
	"sync"
)

var ma sync.Map

func cleanMap() {
	ma.Range(func (k, _ interface{}) bool {
		ma.Delete(k)
		return true
	})
}

func appendMap(k string) {
	v, ok := ma.Load(k)
	x := 0
	if ok {
		x = v.(int)
	} else {
		ma.Store(k, x)
	}
	if x == 0 || x == 3 {
		ma.Delete(k)
		fmt.Println("delete key from map")
		x = 0
	}
	ma.Store(k, x + 1)
}

func main() {


	for i := 0; i < 30; i++ {
		appendMap("a")
	}
	// output vvv： a : 3
	ma.Range(func (k,v interface{}) bool {
		fmt.Printf("vvv： %s : %d \n", k, v)
		return true
	})
	cleanMap()
	// non output
	ma.Range(func (k,v interface{}) bool {
		fmt.Printf("vvv： %s : %d \n", k, v)
		return true
	})
}
