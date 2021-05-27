package main

import "testing"

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(10)
	}
}

func TestJson(t *testing.T) {
	MapJson()
}