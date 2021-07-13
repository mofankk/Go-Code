package grpc

import (
	"fmt"
	"goLearnCode/microservice/message"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"time"
)

func StartStreamRpcClient() {
	StreamRpcClient()
}
func StreamRpcClient() {
	time.Sleep(time.Second)
	conn, err := grpc.Dial("127.0.0.1:8083", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	fmt.Println("客户端启动")
	orderServiceClient := message.NewStreamOrderServiceClient(conn)
	fmt.Println("客户端请求RPC调用，双向流模式")
	orderIDs := []string{"202103122145", "201712120005", "201907061600"}

	// 创建一个管理数据流发送/接收的客户端
	orderInfoClient, err := orderServiceClient.GetOrderInfoStream(context.Background())
	if err != nil {
		panic(err.Error())
	}
	for _,id := range orderIDs {
		orderRequest := message.OrderRequest{OrderId: id}
		err := orderInfoClient.Send(&orderRequest)
		if err != nil {
			panic(err.Error())
		}
	}
	// 关闭发送数据流
	orderInfoClient.CloseSend()
	// 开始就收服务返回的数据流
	for {
		orderInfo, err := orderInfoClient.Recv()
		if err == io.EOF {
			fmt.Println("读取结束")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		if orderInfo.GetOrderId() != "" {
			fmt.Println("读取到的信息: ", orderInfo)
		}
	}
	fmt.Println("客户端关闭")
}