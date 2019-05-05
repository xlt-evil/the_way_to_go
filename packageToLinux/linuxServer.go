package main

import (
	"fmt"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Printf("启动失败！")
		return
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("bug")
		}
		helloWorld(conn)
	}
}

func helloWorld(conn net.Conn) {
	fmt.Println("连接成功")
	fmt.Println(conn.RemoteAddr().String())
	fmt.Fprint(conn, "helloworld")
	conn.Close()
}
