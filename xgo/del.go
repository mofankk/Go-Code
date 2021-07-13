package main

import "fmt"

func main() {

    var a interface{}

    var b int = 12

    a = b

    fmt.Printf("%T +  %v \n", a, a)
    fmt.Printf("%T  + %v \n", b, b)
}
