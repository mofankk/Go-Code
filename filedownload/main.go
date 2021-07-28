package main

import (
	"filedownload/handler"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	route := mux.NewRouter()

	fd := handler.DownloadFile{}
	route.HandleFunc("/download", fd.Download).Methods("GET")
	route.HandleFunc("/muti/download", fd.MutiThreadDownload).Methods("POST")

	server := http.Server{
		Addr: "0.0.0.0:8000",
		Handler: route,
	}
	fmt.Println("服务已启动，端口：8000")
	server.ListenAndServe()
}