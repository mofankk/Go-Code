package modules

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	// 不引入下面这个会报: sql: unknown driver "postgres" (forgotten import?)
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var conn *gorm.DB

// SetUp 创建数据库连接
func SetUp() {

	dsn := "postgres://" +
		"mofan"+ ":" +
		"nihao2021" + "@" +
		"localhost" + "/" +
		"gorm_test" + "?sslmode=disable"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		fmt.Printf("数据库连接失败: %v", err)
		db = nil
		return
	}
	sqlDB := db.DB()
	sqlDB.SetMaxOpenConns(300)
	sqlDB.SetMaxIdleConns(20)

	//fmt.Println(sqlDB.Stats())
	conn = db
}

func GetDB() (*gorm.DB, error) {
	if conn != nil {
		return conn, nil
	}

	return nil, errors.New("数据库连接失败")
}
