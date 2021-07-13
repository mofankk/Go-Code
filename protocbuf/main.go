package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"protobuf/api"
)

func main() {

	book := &api.AddressBook{}

	p1 := &api.Person_PhoneNumber{
		Number: "11111",
		Type:   api.Person_MOBILE,
	}
	var phone []*api.Person_PhoneNumber
	phone = append(phone, p1)

	p := &api.Person{
		Name:        "Mofan",
		Id:          1234,
		Email:       "m@gmail.com",
		Phones:      phone,
		LastUpdated: nil,
	}
	book.People = append(book.People, p)

	// Writing a Message
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	fname := "file/protoc.txt"
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Reading a Message
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	bk := &api.AddressBook{}
	if err := proto.Unmarshal(in, bk); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Println(bk)
}
