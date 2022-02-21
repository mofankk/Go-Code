package main

import "fmt"

// 3. 翻转字符串
func main() {
	fmt.Println(reverseString("vim-go"))
}

func reverseString(s string) (string, bool) {
	sl := len(s)
	if sl > 5000 {
		return s, false
	}
	str := []rune(s)

	for i := 0; i < sl/2; i++ {
		str[i], str[sl-1-i] = str[sl-1-i], str[i]
	}
	return string(str), true
}
