package handler

import (
	"encoding/json"
	"fmt"
	"oldgorm/modules"
	"os"
	"time"
)

func GetAccount(f *os.File) {
	db, err := modules.GetDB()
	if err != nil {
		fmt.Println("DB获取错误")
		return
	}
	//if db != nil {
	//	defer db.Close()
	//} else {
	//	fmt.Println("DB是空的")
	//}
	var info []modules.Account

	//db.LogMode(true)
	s, _ := json.Marshal(db.DB().Stats())
	f.WriteString(string(s) + "\n\n")
	err = db.Where("username like '%" + "x" + "%'").Find(&info).Error
	if err != nil {
		fmt.Printf("查询失败: %v", err)
	}

	//fmt.Println(len(info))
	time.Sleep(1 * time.Second)
}


func GetAccountB(f *os.File) {
	db, err := modules.GetConn()
	if err != nil {
		fmt.Println("DB获取错误")
		return
	}
	//if db != nil {
	//	defer db.Close()
	//} else {
	//	fmt.Println("DB是空的")
	//}
	var info []modules.Account

	db.LogMode(true)
	s, _ := json.Marshal(db.DB().Stats())
	f.WriteString(string(s) + "\n")
	err = db.Where("username like '%" + "x" + "%'").Find(&info).Error
	if err != nil {
		fmt.Printf("查询失败: %v", err)
	}

	//fmt.Println(len(info))
}