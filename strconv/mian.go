package main

import (
	"fmt"
	"strconv"
)

type student struct {
	age int
}

//字符串转换
func main() {
	str := make([]byte, 0, 100) //make用与创建切片的
	//append将类型转换为相应的字符串后添加到byte
	str = strconv.AppendInt(str, 4567, 10) //返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字。
	fmt.Println(string(str))
	str = strconv.AppendBool(str, false)
	fmt.Println(string(str))
	//Format系列把其他类型转换成string类型
	a := strconv.FormatBool(false)
	var ss float64
	ss = 123.23
	b := strconv.FormatFloat(ss, 'g', 12, 64)
	fmt.Println("hahah", b)
	c := strconv.FormatInt(1234, 2)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023) //原样还原
	fmt.Println(a, b, c, d, e)
	//parse把字符串还原成其他类型
	l, _ := strconv.ParseBool("false")
	z, _ := strconv.ParseFloat("123.23", 64)
	x, _ := strconv.ParseInt("1234", 10, 64)
	v, _ := strconv.ParseUint("12345", 10, 64)
	m, _ := strconv.Atoi("2134")
	fmt.Println(l, z, x, v, m)
}
