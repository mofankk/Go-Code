package main

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	Name string
	Age  int
}

func main() {
	var stu Stu
	b, _ := json.Marshal(Stu{
		Name: "hello",
		Age:  12,
	})

	err := json.Unmarshal(b, &stu)
	if err != nil {
		panic(err)
	}

	fmt.Println(stu.Name, stu.Age)
}
