package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//实现linux wget 命令
func main() {
	wget()
}

// os.Args 0 是文件名
func wget() {
	httpurl := ""
	if len(os.Args) > 1 {
		httpurl = os.Args[1]
	} else {
		fmt.Println("wget v.1.0")
		fmt.Println("    by hxy")
		return
	}
	resp, err := http.Get(httpurl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fileLength := resp.ContentLength
	f, err := os.Create(httpurl[strings.LastIndex(httpurl, "/")+1:])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	PrintSpeedOfProgess(fileLength, f, resp)
}

func PrintSpeedOfProgess(length int64, file *os.File, resp *http.Response) {
	buffer := make([]byte, 8*1024)
	r := resp.Body
	basesize := length / 100
	defer func() {
		file.Close()
		r.Close()
	}()
	if basesize == 0 {
		fmt.Printf("this is not a file url ")
		os.Remove(file.Name())
		return
	}
	sum := 0
	block := 0
	for {
		size, err := r.Read(buffer)
		if err != nil {
			if err != io.EOF {
				os.Remove(file.Name())
				fmt.Println(err.Error())
				return
			}
		}
		_, err = file.Write(buffer[:size])
		if err != nil {
			return
		}
		sum += size
		b := int64(sum) / basesize
		if int64(block) < b {
			cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
			cmd.Stdout = os.Stdout
			err = cmd.Run()
			fmt.Print(block, "% [")
			block = 0
			for int64(block) < b {
				fmt.Print("#")
				block++
			}
			str := "%" + strconv.Itoa(101-block) + "s"
			fmt.Printf(str, "]")
		}
		if block == 100 {
			return
		}
	}
}
