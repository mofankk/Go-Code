package handler

import (
	"encoding/json"
	"fmt"
	"oldgorm/modules"
	"time"
)

func GetAccount() {
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

	db.LogMode(true)
	s, _ := json.Marshal(db.DB().Stats())
	fmt.Println(string(s))
	err = db.Where("username like '%" + "x" + "%'").Find(&info).Error
	if err != nil {
		fmt.Printf("查询失败: %v", err)
	}

	fmt.Println(len(info))
	time.Sleep(1 * time.Second)
}


func GetAccountB() {
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
	fmt.Println(string(s))
	err = db.Where("username like '%" + "x" + "%'").Find(&info).Error
	if err != nil {
		fmt.Printf("查询失败: %v", err)
	}

	fmt.Println(len(info))
}