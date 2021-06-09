package handler

import (
	"fmt"
	"oldgorm/modules"
)

func GetAccount() {
	db, err := modules.GetDB()
	if err != nil {
		fmt.Println("DB获取错误")
		return
	}
	if db != nil {
		defer db.Close()
	} else {
		fmt.Println("DB是空的")
	}
	var info []modules.Account

	err = db.Find(&info).Error
	if err != nil {
		fmt.Printf("查询失败: %v", err)
	}

	fmt.Println(len(info))
}
