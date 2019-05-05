package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Student struct {
	Name string
}

func main() {
	client, err := rpc.Dial("tcp", ":9090")
	if err != nil {
		log.Fatal(err.Error())
	}
	var name string
	stu := &Student{""}
	name = "test"
	err = client.Call("MyStudent.SetName", &name, stu)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(stu.Name, name)
	s := &Student{"我是大笨蛋"}
	client.Call("MyStudent.GetName", s, &name)
	fmt.Println(name)
}
