package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	tutorial "play.init/src"
)

func main() {
	t := tutorial.Person{
		Name:        "k",
		Id:          3,
		LastUpdated: &timestamppb.Timestamp{Seconds: 300, Nanos: 30},
		Username:    "k@k",
		Phones:      []*tutorial.Person_PhoneNumber{{Number: "1212323", Type: tutorial.Person_WORK}, {Number: "12352222222", Type: tutorial.Person_HOME}},
	}
	fmt.Println(t)
	fmt.Println(t.GetUsername())

	t.Username = "asdasds"

	address := tutorial.AddressBook{
		People: []*tutorial.Person{&t, &t},
	}

	fmt.Println(address)

	_, errdoWriteFile := doWriteFile("text.out", &address)
	if errdoWriteFile != nil {
		log.Fatalln("doWriteFile", errdoWriteFile)
		return
	}
	_, errdoReadFile := doReadFile("text.out")
	if errdoReadFile != nil {
		log.Fatalln("doReadFile", errdoReadFile)
		return
	}

}

func doWriteFile(file string, t proto.Message) (bool, error) {
	data, err := proto.Marshal(t)
	if err != nil {
		log.Fatalln("Marshal fail")
		return false, err
	}

	err2 := ioutil.WriteFile(file, data, 0644)
	if err2 != nil {
		log.Fatalln("Writing fail")
		return false, err2
	}

	fmt.Println("File Writted")
	return true, nil
}

func doReadFile(file string) (bool, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln("ReadFile")
		return false, err
	}
	t := tutorial.Person{}
	err2 := proto.Unmarshal(data, &t)
	if err2 != nil {
		log.Fatalln("Unmarshal")
		return false, err2
	}
	fmt.Println("File read")
	fmt.Println(t)
	return true, nil
}
