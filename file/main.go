package main

import (
	"fmt"
	"os"
)

var dirname []string
var filename []string
var count int

func main() {

	ReadInfo("c:/")
	fmt.Println(dirname)
	fmt.Println(filename)
	fmt.Println(count)

}

func ReadInfo(src string) {
	f, _ := os.Open(src)
	f2, _ := f.Readdir(-1)
	for i, _ := range f2 {
		if f2[i].IsDir() {
			dirname = append(dirname, f2[i].Name())
			ReadInfo(src + "/" + f2[i].Name())
		} else {
			filename = append(filename, f2[i].Name()+"\n")
		}
		count++
	}
	return
}
