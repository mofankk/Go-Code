package transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"microservice1/endpoint"
	"net/http"
	"strconv"
)

func DecodeUserRequest (ctx context.Context, r *http.Request) (interface{}, error) {
	//uid := r.URL.Query().Get("id")
	vars := mux.Vars(r)
	uid := vars["id"]
	if uid != "" {
		id, err := strconv.Atoi(uid)
		if err != nil {
			return nil, err
		}
		return endpoint.UserRequest{Id: id, Method: r.Method}, nil
	}
	return nil, errors.New("参数错误")
}

func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, res interface{}) error {
	w.Header().Set("Content-type", "application/json")

	return json.NewEncoder(w).Encode(res)
}
