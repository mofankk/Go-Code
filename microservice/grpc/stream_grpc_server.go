package grpc

import (
	"fmt"
	"goLearnCode/microservice/message"
	"google.golang.org/grpc"
	"io"
	"net"
	"time"
)

func (o *OrderServiceImpl) GetOrderInfoStream(stream message.StreamOrderService_GetOrderInfoStreamServer) error {
	for {
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("数据读取结束")
			return err
		}
		if err != nil {
			panic(err.Error())
			return err
		}
		fmt.Printf("请求订单ID为 %s 的数据\n", orderRequest.GetOrderId())
		// 定义好所有的订单信息
		orderMap := map[string]message.OrderInfo{
			"202103122148": {OrderId: "202103122148", OrderName: "Mac Book Pro", OrderStatus: "攒钱中"},
			"201712120005": {OrderId: "201712120005", OrderName: "Xiaomi Book Pro", OrderStatus: "已付款"},
			"201907061600": {OrderId: "201907061600", OrderName: "iPhone", OrderStatus: "已付款"},
		}
		result := orderMap[orderRequest.GetOrderId()]
		if result.OrderId == "" {
			fmt.Println("没有该订单")
		}
		err = stream.Send(&result)
		time.Sleep(time.Second)
		if err == io.EOF {
			fmt.Println(err)
			return err
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func StartStreamGrpcServer() {
	server := grpc.NewServer()
	//注册
	message.RegisterStreamOrderServiceServer(server, new(OrderServiceImpl))

	listen, err := net.Listen("tcp", "127.0.0.1:8083")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(listen)
}