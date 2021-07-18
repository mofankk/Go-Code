package endpoint

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"microservice1/service"
	"net/http"
)

type UserRequest struct {
	Id     int    `json:"id"`
	Method string `json:"method"`
}

type UserResponse struct {
	Data string `json:"data"`
}

func GenUserEndPoint(us service.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (res interface{}, err error) {
		r := request.(UserRequest)
		data := "nothing"
		if r.Method == http.MethodGet {
			data = us.GetName(r.Id)
		} else if r.Method == http.MethodDelete {
			if err = us.DelUser(r.Id); err != nil {
				data = "删除失败"
			}
			data = fmt.Sprintf("删除ID为 %d 的用户成功", r.Id)
		}
		return UserResponse{Data: data}, err
	}
}
