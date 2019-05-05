package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	c, e := net.Dial("tcp", "192.168.88.95:8888")
	if e != nil {
		fmt.Printf("连接失败" + e.Error())
	}
	b, e := ioutil.ReadAll(c)
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println(string(b))
}
