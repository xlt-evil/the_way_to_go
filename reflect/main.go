package main

import (
	"fmt"
	"reflect"
)

//深入了解反射
//什么是反射
//指程序可以访问、检测和修改它本身状态或行为的一种能力//比较专业的说法
//通常我们写一个类型都是希望这个类型能做些什么，而反射就是这个类型能对自己做些什么//面向的事物不同比如我很轻易的看到其他人长什么样的但不能看到自己的我需要哪一个镜子反射出来自己的一些特性
//大量的反射会损失一定的性能

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Printf("my name is %s and  i'm %d", u.Name, u.Age)
}

//利用反射得到这个包的信息
//我只知道有这个结构体其他的一概不知，利用反射告诉我他拥有什么
func Info(o interface{}) {
	//得到类型
	t := reflect.TypeOf(o)
	//输出其类型
	fmt.Println("这个类型是", t.Name())
	//得到类型的值
	v := reflect.ValueOf(o)
	fmt.Println("类型的值", v)
	var num = t.NumField() //可以不写类型自动推断
	fmt.Println("这个类型拥有的参数个数:", num)
	for i := 0; i < num; i++ {
		f := t.Field(i)
		val := v.Field(i).Interface() //返回的值保存在interface里面
		v.Kind()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
	//下面是反射得到类型的方法名称和类型的方法的类型
	var numMenthon = t.NumMethod() //
	fmt.Println("拥有的方法", numMenthon)
	for i := 0; i < numMenthon; i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
	fmt.Println(t.Kind().String()) //kind用与输出判断的值时什么类型
}

func main() {
	user := User{1, "黄希云", 21}
	Info(user)

}
