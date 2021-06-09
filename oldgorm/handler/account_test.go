package handler

import "testing"

func BenchmarkGetAccount(b *testing.B) {
	GetAccount()
}

func BenchmarkGetAccountB(b *testing.B) {
	GetAccountB()
}

