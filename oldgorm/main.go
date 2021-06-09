package main

import (
	"oldgorm/handler"
	"oldgorm/modules"
)

func main() {
	modules.SetUp()
	handler.GetAccount()
}
