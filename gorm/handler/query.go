package handler

import (
	"gorm/module"
)

// 使用数据库连接池
func Query() {
	db, _ := module.GetDB()
	var info []module.Account
	db.Find(&info)
}

// 不使用数据库连接池
func SignQuery() {
	conn := module.GetDBConn()
	var info []module.Account
	conn.Find(&info)
}
