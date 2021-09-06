package main

import (
	"log"
	"strings"
)

func main() {

	str := "ab ba    "
	log.Println(len(str))
	log.Println(len(strings.TrimSpace(str)))
	log.Println(strings.TrimSpace("     abc daa aaa"))
	log.Println(strings.TrimSpace("abcddc ba"))
}
