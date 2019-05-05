package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//ioutil 的一些基本方法展示
func main() {
	path, err := os.Getwd() //获取当前环境运行路径
	fmt.Println(path)
	b, err := ioutil.ReadFile("test.txt") //读文件
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("bug")
	}
	ioutil.WriteFile("123.txt", b, os.FileMode(123)) //写文件
	fmt.Println(string(b))
	name := string(b)
	s := strings.Split(name, "/")
	fmt.Println(s)
	//ioutil.ReadAll()通常去读http返回的东西
}
