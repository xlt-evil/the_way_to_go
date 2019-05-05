package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"` //输出子xml
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIp   string   `xml:"serverIP"`
}

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}

func main() {
	generateXml()
}

func ParseXml() {
	file, err := os.Open("D:/HxyGo/src/the-way-to-go/learn_8/config.xml") //for read access
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	fmt.Println(v.Description)
	fmt.Println(v.Svs)
	fmt.Println(v.Version)
}

func generateXml() {
	name := xml.Name{}
	v := &Servers{Version: "2"}
	v.Svs = append(v.Svs, server{name, "Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{name, "Beijing_VPN", "127.0.0.2"})
	data, err := xml.MarshalIndent(v, " ", " ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	file, _ := os.Create("config1.xml")
	file.Write([]byte(xml.Header))
	file.Write(data)
	file.Close()
}
