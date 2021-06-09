package module

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetDb(t *testing.T) {
	conn, err := GetDB()
	if err == nil {
		fmt.Println("测试通过")
		d, _ := conn.DB()
		b, _ := json.Marshal(d.Stats())
		fmt.Println(string(b))
	} else {
		fmt.Println("测试失败")
	}
}