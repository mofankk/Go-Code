package main

import "encoding/json"


func main() {
	type My {
		Name string
		Age  uint
		Online bool 
	}
	
	my := My {
		Name: "cyp",
		Age: 100,
		Online: true,
	}
	
	b, _ := json.Marshal(my)
	
	var temp My
	
	err := json.NewDecoder()
	
