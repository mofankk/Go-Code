package main

import (
	"fmt"
	db2 "gorm_mysql/db"
	"gorm_mysql/model"
)

func main() {
	fmt.Println("Start")

	//model.InitDB()

	db := db2.DBConn()

	var user []model.User
	db.Preload("Company").Find(&user)

	fmt.Println(user)

	fmt.Println("End")
}
