package main

import (
	"fmt"
	"strings"
)

// 4. 判断两个给定的字符串排序后是否一致
func main() {
	fmt.Println(isRegroup("abc", "cba"))
}

func isRegroup(s1, s2 string) bool {
	sl1 := len(s1)
	sl2 := len(s2)
	if sl1 != sl2 || sl1 > 5000 || sl2 > 5000 {
		return false
	}

	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}

	return true
}
