package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//生成JSON
type Server struct {
	ServerName string `json:"servername"`
	ServerIP   string `json:"ip"`
}
type Serverslice struct {
	Servers []Server
}

type Web struct {
	Web []Config
}

type Config struct {
	Name string
	Url  string
}

func main() {
	//generateJSON()
	//ParseJSON()
	file, err := os.Open("./the-way-to-go/json/config.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	b, _ := ioutil.ReadAll(file)
	fmt.Println(string(b))
	c := Web{}
	json.Unmarshal(b, &c)
	fmt.Println(c)
}

//字段小写无法映射成为json
type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func generateJSON() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

func ParseJSON() {
	var s Student
	s.Age = 10
	s.Name = "hxy"
	fmt.Println(s)
	data, err := json.Marshal(s)
	if err != nil {
		log.Fatal("bug")
	}
	fmt.Println(string(data))
	var str Student
	json.Unmarshal(data, &str)
	fmt.Println(str)
}
