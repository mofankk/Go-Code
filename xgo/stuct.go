package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Hello struct {
	H int
	E string
	L bool
}

type World struct {
	Hello
	W string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(ww http.ResponseWriter, r *http.Request) {

		var w World
		w.H = 1
		w.E = "hello"
		w.L = true
		w.W = "world"

		b, _ := json.Marshal(w)
		ww.Write(b)

	}).Methods("GET")

	src := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8000",
	}

	src.ListenAndServe()

}
