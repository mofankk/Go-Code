package transport

import (
	"context"
	"encoding/json"
	"errors"
	"microservice2/endpoint"
	"net/http"
	"strconv"
)

func GetUserInfo_Request (_ context.Context, request *http.Request, r interface{}) error {
	req := r.(endpoint.UserRequest)
	request.URL.Path += "/user/" + strconv.Itoa(req.Id)

	return nil
}

func GetUserInfoResponse(_ context.Context, res *http.Response) (response interface{}, err error) {

	if res.StatusCode > 400 {
		return nil, errors.New("No Data!!!")
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return
}
