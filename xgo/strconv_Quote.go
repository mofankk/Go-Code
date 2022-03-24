package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := `" Hello Welcome 
to GeeksforGeeks "
`
	fmt.Println(strconv.Quote(str))

}

// result：
// "\" Hello Welcome \nto GeeksforGeeks \"\n"
