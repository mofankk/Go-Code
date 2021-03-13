package grpc

import (
	"testing"
	"time"
)

func TestGrpc(T *testing.T) {
	go StartGrpcServer()
	go StartGrpcClient()
	time.Sleep(time.Second)
}
