package grpc

import (
	"fmt"
	"goLearnCode/microservice/message"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

func StartGrpcClient() {
	GrpcClient()
}

func GrpcClient() {
	conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)

	orderRequest := &message.OrderRequest{OrderId: "201712120005", TimeStamp: time.Now().Unix()}
	orderInfo, err := orderServiceClient.GetOrderInfo(context.Background(), orderRequest)
	if err != nil {
		panic(err.Error())
	}
	if orderInfo != nil {
		fmt.Println(orderInfo.GetOrderId())
		fmt.Println(orderInfo.GetOrderName())
		fmt.Println(orderInfo.GetOrderStatus())
	}
}
