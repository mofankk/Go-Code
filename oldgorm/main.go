package main

import (
	"fmt"
	"time"

	"oldgorm/handler"
)

/*
这个是 V1
GORM V2 moved to https://github.com/go-gorm/gorm

GORM V1 Doc https://v1.gorm.io/
*/

func main() {
	//modules.SetUp()
	for i := 0; i < 1000; i++ {
		go handler.GetAccount()
	}

	fmt.Println("最后一个")
	handler.GetAccountB()

	//go handler.GetAccountB()
	//go handler.GetAccountB()
	//go handler.GetAccountB()
	//go handler.GetAccountB()
	//go handler.GetAccountB()

	time.Sleep(10 * time.Second)

}
