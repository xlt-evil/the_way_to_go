package main

import (
	"fmt"
	"strings"
)

//常用字符串处理
func main() {
	var s string = "我叫黄希云，我叫一歌笑byevil"
	//字符串s是否包含substr
	fmt.Println(strings.Contains(s, "黄希云"))
	s1 := []string{"foo", "bar", "baz"}
	//字符串链接，把slice a 通过seq连接
	s2 := strings.Join(s1, ",")
	fmt.Println(s2)
	//找到对用的第一个下标数，美没有为-1
	s3 := strings.Index(s, "我")
	fmt.Println(s3)
	//repeat 重复字符串count次
	fmt.Println(strings.Repeat(s, 2))
	//把字符串中把old字符串替换为new字符串n代表替换次数 小于 0 代表全替换
	fmt.Println(strings.Replace(s, "黄希云", "大佬", 2))
	//把字符串按照seq分割返回slice
	fmt.Println(strings.Split(s, ","))
	//在s中掉指定的字符
	ss := "12334  51我"
	fmt.Println(strings.Trim(ss, "我"))
	//去掉空格符并且按照slice返回
	fmt.Println(strings.Fields(ss))
}
