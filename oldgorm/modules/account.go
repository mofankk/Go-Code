package modules

type Account struct {
	AccountId        uint `gorm:"primary_key;AUTO_INCREMENT" `
	Username         string
	Password         string
	Role             int `gorm:"default:1"`
	Memo             string
	AccountStatus    int `gorm:"default:0"`
	AccountType      int `gorm:"default:1"`
	AccountTerritory string
}

func (Account) TableName() string {
	return "account"
}
