package main

import (
	"fmt"
	"log"
	"os"
)

//了解os包的基本用法
func main() {
	os.Mkdir("test", 111)              //创建文件夹
	os.MkdirAll("test/test/test", 123) //创建多级文件夹
	os.RemoveAll("test")               //删除多级文件夹
	os.Remove("test/test/test")        //定位要明确，如果删除的文件夹下有文件夹就无法删除
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(file)
	}
	b := make([]byte, 1024*8) //读的大小8kb通常是读写最快的时候
	n, err := file.Read(b)    //读取文件返回字符
	fmt.Println(string(b), n)
	f, err := os.Create("go.txt") //创建文件
	os.NewFile(1, "goto.txt")     //新建一个文件不保存的=
	f.WriteString("hahaha")       //写入文件
	f.Write(b)                    //以[]byte方式写入
	fmt.Println(f.Name())
	f.Sync() //Sync递交文件的当前内容进行稳定的存储。一般来说，这表示将文件系统的最近写入的数据在内存中的拷贝刷新到硬盘中稳定保存
	f.Close()
	file.Close()
	os.Exit(1) //结束程序运行以code码返回
}
