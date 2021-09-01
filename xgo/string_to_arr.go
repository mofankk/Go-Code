package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//func stringToArr(st string) interface{} {
//	stt := st[1:]
//	length := len(stt)
//	stx := stt[0 : length-1]
//	str := "[" + stx + "]"
//	fmt.Println(str)
//	str = strings.Replace(str, "(", "[", -1)
//	str = strings.Replace(str, ")", "]", -1)
//
//	fmt.Println(str)
//	var p [][]interface{}
//	b, _ := json.Marshal(str)
//	err := json.Unmarshal(b, &p)
//	if err != nil {
//		fmt.Printf("xxx:%v", err)
//	}
//	return p
//}

func stringToArr(str string) [][]int64 {
	strArr := strings.FieldsFunc(str, func(r rune) bool {
		if unicode.IsNumber(r) {
			return false
		}
		return true
	})
	fmt.Println(strArr)
	resArr := make([][]int64, len(strArr)/2)
	for i := range strArr {
		x, y := i/2, i%2
		fmt.Println(i)
		if y == 0 {
			resArr[x] = make([]int64, 2)
			resArr[x][0], _ = strconv.ParseInt(strArr[i], 10, 64)
		} else {
			resArr[x][1], _ = strconv.ParseInt(strArr[i], 10, 64)
		}
	}
	return resArr
}

func main() {
	str := "((1,2),(3,4),(5,6))"
	//p := stringToArr(str)
	//p := strings.FieldsFunc(str, func(r rune) bool {
	//	if unicode.IsNumber(r) {
	//		return false
	//	}
	//	return true
	//})
	//fmt.Printf("%v", p)

	fmt.Println(stringToArr(str))
}
