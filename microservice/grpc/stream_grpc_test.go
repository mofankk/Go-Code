package grpc

import (
	"testing"
	"time"
)

func TestSteamGpc(T *testing.T) {
	go StartStreamGrpcServer()
	go StartStreamRpcClient()
	time.Sleep(10 * time.Second)
}
