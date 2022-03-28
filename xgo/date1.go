package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	//fmt.Printf("Go launched at %s\n", t.Local())

	t1 := "2022-03-04 15:04:05"
	tf, err := time.Parse("2006-01-02 15:04:05", t1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tf: ", tf.Local())

	tn := time.Now()
	tns := tn.Local()
	fmt.Println("tns: ", tns)
}
