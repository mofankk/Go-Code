package main

import "fmt"

type DB struct {
	A int
}

func GetDB() (DB, int) {
	db := DB{}
	a := 10

	fmt.Println(&db, &a)
	return db, a

}

func main() {
	x, b := GetDB()

	fmt.Println(&x, &b)

}
