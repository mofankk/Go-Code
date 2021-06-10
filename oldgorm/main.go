package main

import (
	"os"
	"strconv"
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
	f, _ := os.Create("db.log")
	defer f.Close()
	for i := 0; i < 1000; i++ {
		f.WriteString(strconv.Itoa(i) + ", ")
		go handler.GetAccount(f)
	}

	f.WriteString("最后一个\n\n")
	handler.GetAccount(f)
	//handler.GetAccountB(f)

	//go handler.GetAccountB()
	//go handler.GetAccountB()
	//go handler.GetAccountB()
	//go handler.GetAccountB()
	//go handler.GetAccountB()

	time.Sleep(10 * time.Second)

}
