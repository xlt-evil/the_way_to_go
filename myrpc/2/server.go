package main

import (
	"log"
	"net"
	"net/rpc"
)

type Student struct {
	Name string
}

type MyStudent Student

func (t *MyStudent) GetName(stu *Student, name *string) error {
	*name = stu.Name
	return nil
}

func (t *MyStudent) SetName(name *string, stu *Student) error {
	stu.Name = *name
	return nil
}

func main() {
	arith := new(MyStudent)
	err := rpc.Register(arith)
	if err != nil {
		log.Fatal(err.Error())
	}
	l, err := net.Listen("tcp", ":9090")
	for {
		conn, _ := l.Accept()
		rpc.ServeConn(conn)
	}
}
