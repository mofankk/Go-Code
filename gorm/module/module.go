package module

import (
	"time"
)

// Food
// 字段注释说明了gorm库把struct字段转换为表字段名长什么样子。
// 字段定义后面使用两个反引号``包裹起来的字符串部分叫做标签定义，这个是golang的基础语法，不同的库会定义不同的标签，有不同的含义,比如json库
type Food struct {
	Id     int     `gorm:"primaryKey"`	//表字段名为：id
	Name   string  //表字段名为：name
	Price  float64 //表字段名为：price
	TypeId int     //表字段名为：type_id

	CreateTime time.Time `gorm:"default null"`
}

// TableName 设置表名，这里属于自定义表名
func (Food) TableName() string {
	return "food"
}



func UpdateDB() {

	//db := GetDB()

}


