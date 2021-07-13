package main

import (
	"fmt"
	"os"
)

func main() {

	err := os.Remove("abcdefg.txt")
	if err != nil {
		fmt.Println("文件删除失败")
	} else {
		fmt.Println("文件删除成功")
	}

}

