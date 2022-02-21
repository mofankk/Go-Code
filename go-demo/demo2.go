package main

import (
	"fmt"
	"strings"
)

// 2. 判断字符串中字符是否全都不同
func main() {
	a := "hello,world"
	fmt.Println(isUniqueString(a)) // true
	fmt.Println(len(""))           // 0
}

// use strings.Count()
func isUniqueString(s string) bool {
	// If substr is an empty string, Count returns 1 + the number of Unicode code points in s. so, this is 126
	if strings.Count(s, "") > 126 {
		return false
	}

	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}

	return true
}

// use strings.LastIndex() or strings.Index()
func isUniqueString2(s string) bool {
	if strings.Count(s, "") > 127 {
		return false
	}

	for i, v := range s {
		if v > 127 {
			return false
		}

		if strings.LastIndex(s, string(v)) != i {
			return false
		}
	}

	return true
}
