package main

import (
	"fmt"
	"reflect"
)

//go 的tag都是使用在反射中

type User struct {
	Name   string "user name"
	Passwd string "user password"
}

type S struct {
	F string `hxy:"666" color:"123"`
}

func main() {
	B()
}

func A() {
	user := &User{"chronos", "pass"}
	s := reflect.TypeOf(user).Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag) //将tag输出出来
	}
}

func B() {
	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("hxy"))
}
