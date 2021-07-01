package api

import (
	"encoding/json"
	"net/http"
)

type Hello struct {

}

func (*Hello) SayHello(w http.ResponseWriter, r *http.Request) {

	type Res struct {
		Name string 	`json:"name"`
		World string 	`json:"world"`
	}

	res := Res{Name: "mofan", World: "Hello, World"}

	b, _ := json.Marshal(res)

	w.Write(b)
}
