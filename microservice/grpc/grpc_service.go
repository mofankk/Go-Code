package grpc

import (
	"errors"
	"fmt"
	"goLearnCode/microservice/message"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"time"
)

type OrderServiceImpl struct {

}

func StartGrpcServer() {
	server := grpc.NewServer()
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))

	listen, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(listen)
}

//获取订单信息
func (o *OrderServiceImpl) GetOrderInfo(ctx context.Context, request *message.OrderRequest) (*message.OrderInfo, error) {
	orderMap := map[string]message.OrderInfo{
		"202103122148": {OrderId: "202103122148", OrderName: "Mac Book Pro", OrderStatus: "攒钱中"},
		"201712120005": {OrderId: "201712120005", OrderName: "Xiaomi Book Pro", OrderStatus: "已付款"},
		"201907061600": {OrderId: "201907061600", OrderName: "iPhone", OrderStatus: "已付款"},
	}
	var response message.OrderInfo
	current := time.Now().Unix()
	if request.TimeStamp > current {
		response = message.OrderInfo{OrderId: "", OrderName: "", OrderStatus: "订单信息异常"}
		return &response, nil
	}
	response = orderMap[request.GetOrderId()]
	if response.GetOrderId() != "" {
		fmt.Println("订单信息查找成功...")
		return &response, nil
	} else {
		return nil, errors.New("Server Error")
	}
}