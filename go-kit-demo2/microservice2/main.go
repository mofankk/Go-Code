package main

import (
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
	"net/url"
	"context"
	"microservice2/transport"
	"microservice2/endpoint"
	"os"
)

func main() {
	tgt, err := url.Parse("http://192.168.3.153:8080")
	if err != nil {
		log.Fatal(err)
	}

	client := httptransport.NewClient(http.MethodGet, tgt, transport.GetUserInfo_Request, transport.GetUserInfoResponse)

	getUserInfo := client.Endpoint()

	ctx := context.Background()

	res, err := getUserInfo(ctx, endpoint.UserRequest{Id: 101})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	userInfo := res.(endpoint.UserResponse)
	fmt.Println(userInfo.Data)
	fmt.Println(res)
}