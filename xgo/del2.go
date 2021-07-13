package main

import "fmt"

func main() {
    var a, sum int = 1, 0
    for i := 0; i <= 30; i++ {
        sum += a
        a *= 2
    }
    fmt.Printf("最大值: %d\n", sum)
    sum++
    fmt.Printf("最小值: %d", sum)
}
