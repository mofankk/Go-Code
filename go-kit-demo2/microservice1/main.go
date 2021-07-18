package main

import (
	"fmt"
	transporthttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"log"
	"microservice1/endpoint"
	"microservice1/service"
	"microservice1/transport"
	"microservice1/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	us := service.UserService{}
	endPoint := endpoint.GenUserEndPoint(us)

	serviceHandler := transporthttp.NewServer(endPoint, transport.DecodeUserRequest, transport.EncodeUserResponse)

	router := mux.NewRouter()
	router.Handle("/user/{id}", serviceHandler).Methods(http.MethodGet, http.MethodDelete)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.Write([]byte(`{"Status": "ok"}`))
	}).Methods(http.MethodGet)


	// 服务退出时，优雅地反注册服务
	errChan := make(chan error)
	go func() {
		// 注册服务
		utils.RegisterService()
		err := http.ListenAndServe(":8080", router)
		if err != nil {
			errChan <- err
		}
	}()
	go func() {
		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sigChan)
	}()

	getErr := <-errChan
	utils.DeRegisterService()
	log.Println(getErr)
}