package main

import (
	"os"
	"fmt"
)

func main() {
	env := os.Getenv("CURR_ENV")
	fmt.Printf("%s\n", env)
}
