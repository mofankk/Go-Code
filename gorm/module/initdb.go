package module

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDB 创建数据库连接
func GetDB() *gorm.DB{

	dsn := "user=mofan password=nihao2021 dbname=gorm_test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("有内鬼，终止交易！")
	}
	fmt.Println("数据库连接成功")
	return db
}