package handler

import (
	"testing"
)


func BenchmarkQuery(b *testing.B) {
	Query()
}

func BenchmarkSignQuery(b *testing.B) {
	SignQuery()
}