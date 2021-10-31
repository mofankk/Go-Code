package main

import (
	"os"
	"fmt"
)

func main() {
	//name := os.Getenv("NAME")
	name := GetenvDefault("NAME", "Mofankk")
	fmt.Println(name)
}

func GetenvDefault(key, defVal string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defVal
}
