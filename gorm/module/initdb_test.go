package module

import (
	"fmt"
	"testing"
)

func TestGetDb(t *testing.T) {
	db := GetDB()
	if db != nil {
		fmt.Println("测试通过")
	} else {
		fmt.Println("测试失败")
	}
}
