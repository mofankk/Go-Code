package main

import (
	"os"
	"fmt"
)

func main() {
	env := os.Getenv("CURR_ENV")
	fmt.Printf("%s\n", env)
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
