package main

import (
	"casbin/modules"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func main() {
	db := modules.GetDB()
	a, _ := gormadapter.NewAdapterByDB(db)
	e, _ := casbin.NewEnforcer("path/model.conf", a)
	err := e.LoadPolicy()
	if err != nil {
		fmt.Println("规则载入失败")
	}
	ok, err := e.AddRoleForUser("alice", "data2_admin")
	if ok {
		fmt.Println("为用户添加角色成功")
	}
	ok, err = e.AddPolicy("data2_admin", "data1", "read")
	if ok {
		fmt.Println("添加角色对应资源访问成功")
	}

	ok, err = e.Enforce("alice", "data1", "read")
	if ok {
		fmt.Println("可以访问")
	} else {
		fmt.Println("不可以访问")
	}

}