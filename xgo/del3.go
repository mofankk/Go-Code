package main

import "fmt"

func main() {

    type X struct {
        a int
        b string
    }

    x := X{}
    x.a = 1
    x.b = "hello"

    var y interface{}

    y = x

    fmt.Printf("%T +  %v \n", x, x)
    fmt.Printf("%T  + %v \n", y, y)
}
