package module

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// SetUp 创建数据库连接
func init() {

	dsn := "user=mofan password=nihao2021 dbname=gorm_test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("有内鬼，终止交易！")
	}
	fmt.Println("数据库连接成功")

	sqlDB, err := conn.DB()
	if err != nil {
		db = nil
		return
	}
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetMaxIdleConns(10)

	db = conn
}

func GetDB() (*gorm.DB, error) {
	if db == nil {
		return nil, errors.New("数据库连接失败")
	}
	return db, nil
}

func GetDBConn() *gorm.DB {
	dsn := "user=mofan password=nihao2021 dbname=gorm_test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("有内鬼，终止交易！")
	}
	return conn
}