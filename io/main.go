package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

//了解io复制文件
func main() {
	//target := "3.jpg"
	//source := "1.jpg"
	//myCopy(source,target)
	sss := "2018-05-20"
	now := time.Now()
	nows, _ := time.Parse("2006-01-02", sss)
	fmt.Println(now, nows)
	b := now.Before(nows)
	fmt.Println(b)

}

func copyFile(source, target string) {
	start := time.Now()
	file, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	files, err := os.Create(target)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer files.Close()
	io.Copy(files, file)
	end := time.Now()
	fmt.Println(end.Sub(start)) //3.0083
}

func myCopy(source, target string) {
	start := time.Now()
	file, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	files, err := os.Create(target)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer files.Close()
	b := make([]byte, 8*1024)
	for {
		_, err := file.Read(b)
		if err != nil {
			break
		}
		files.Write(b)
	}
	end := time.Now()
	fmt.Println(end.Sub(start)) //2.0048
}
