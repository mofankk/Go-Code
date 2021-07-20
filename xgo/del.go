package main

import "fmt"

func main() {

    var a interface{}

    var b int = 12

    a = b


    aaa := getstr("hello")

    fmt.Printf("%T +  %v \n", a, a)
    fmt.Printf("%T  + %v \n", b, b)
}


func getstr(str string) string {
	return "hello"
}
