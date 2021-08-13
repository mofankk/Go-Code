package model

// `User` belongs to `Company`, `CompanyID` is the foreign key
type User struct {
	Id      int `gorm:"primaryKey;autoincrement"`
	Name    string
	Sn      int
	Company Company `gorm:"foreignKey: Sn; references: Sn"`
}

func (User) TableName() string {
	return "user"
}

type Company struct {
	ID   int
	Sn   int
	Name string
}

func (Company) TableName() string {
	return "company"
}

//func InitDB() {
//	db := db2.DBConn()
//	db.AutoMigrate(&User{})
//	db.AutoMigrate(&Company{})
//}
