package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 5. 字符串替换问题
func main() {
	fmt.Println(replaceBlank("vim go  "))
}

// use strings.Replace()
// unicode.IsLetter()
func replaceBlank(s string) (string, bool) {
	if len(s) > 1000 {
		return s, false
	}

	for _, v := range s {
		if string(v) != " " && unicode.IsLetter(v) == false {
			return s, false
		}
	}
	return strings.Replace(s, " ", "%20", -1), true
}
