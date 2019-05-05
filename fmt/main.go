package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//in_1()
	//in_2()
	//in_3()
	//in_4()
	//in_5()
	b := make([]byte, 10)
	s := b[0:2:3] //在切片的同时指定容量
	fmt.Println(cap(s))
}

//scanf 输入 按指定格式赋值
func in_1() {
	var str string
	var i int
	var b bool
	fmt.Scanf("%s %d %t", &str, &i, &b)
	fmt.Println(str, i, b)
}

//按顺序赋值
func in_2() {
	var str string
	var i int
	fmt.Scan(&str, &i)
	fmt.Println(str, i)
}

//利用bufio的赋值
func in_3() {
	var inputReader *bufio.Reader
	var input string
	var err error
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("please enter some input :")
	input, err = inputReader.ReadString('s') //读到s终止
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}

func in_4() {
	//var str string
	//var b [1024]byte
	//os.Stdin.Read(b[:])
	//fmt.Println(b)
	//str = string(b[:])
	//fmt.Println(str)
	var buffer [512]byte
	n, err := os.Stdin.Read(buffer[:]) //读取一字节？
	if err != nil {

		fmt.Println("read error:", err)
		return
	}
	fmt.Println("count:", n, ", msg:", string(buffer[:]))
}

//命令行参数输入
func in_5() {
	var str string
	var i int
	str = os.Args[1]
	i, _ = strconv.Atoi(os.Args[2])
	fmt.Println(str, i)
}
