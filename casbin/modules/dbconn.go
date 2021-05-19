package modules

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDB 创建数据库连接
func GetDB() *gorm.DB{

	dsn := "user=robot password=robot host=192.168.3.73 dbname=robot_api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("有内鬼，终止交易！")
	}

	return db
}
