package main

import (
	pb "github.com/Wayne-Liu/hello-grpc/examples/protoct_tutorials"
	"github.com/golang/protobuf/proto"
	"log"
	"io/ioutil"
)

func main()  {
	p := pb.Person{
		Id: 12354,
		Name: "Wayne Liu",
		Email: "liu.wp@inpsur.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number:"5552-33215",Type:pb.Person_HOME},
		},
	}

	book := &pb.AddressBook{}
	book.People = []*pb.Person{&p}
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	//fname,err := os.OpenFile("test.txt",0,755)
	fname := "test.txt"
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book2 := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book2); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	println(book2.People[0].Name)
	println(book2.People[0].Phones[0].Number)

}

