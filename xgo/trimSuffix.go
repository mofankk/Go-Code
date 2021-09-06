package main

import (
	"log"
	"strings"
)

func main() {
	log.Println(strings.TrimSuffix("abba", "ba"))
	log.Println(strings.TrimSuffix("abcdaaaaa", "abcd"))
	log.Println(strings.TrimSuffix("abcddcba", "dcba"))
}
