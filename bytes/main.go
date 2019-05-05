//bytes包实现了操作[]byte的常用函数。本包的函数和strings包的函数相当类似。
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	s := []byte{1, 2, 4, 5, 7, 8, 8}
	s1 := []byte{1, 2, 4, 5, 6, 9, 8}
	n := bytes.Compare(s, s1) //比较两个[]byte 按字典序比较 相等是0，小于是-1 大于是1 比较到出现一个不一样的byte为止返回
	fmt.Println(n)

	b := bytes.Equal(s, s1) //比较两个[]byte是否相等
	fmt.Println(b)

	//判断两个utf-8编码切片
	b = bytes.EqualFold(s, s1)
	fmt.Println(b)

	s = []byte{65, 66, 67, 66}
	r := bytes.Runes(s) //返回等价的runes//byte 与rune 的共同点是 在一个字节内的码点是一样的操出范围就有可能需要几个byte表示utf-8 但是rune依旧只要一个
	fmt.Println(string(r))

	b = bytes.HasPrefix(s, []byte{65, 66}) //看下是否有前缀(首位开始)切片prefix
	fmt.Println(b)

	b = bytes.HasSuffix(s, []byte{66, 67}) //看下是否有后缀(到尾部结束)
	fmt.Println(b)

	b = bytes.Contains(s, []byte{66}) //判断是否包含子序列
	fmt.Println(b)

	n = bytes.Count(s, []byte{66}) //看下包含几个不重叠的统计
	fmt.Println(n)

	n = bytes.Index(s, []byte{66}) //出现第一个子切片的位置
	fmt.Println(n)

	s2 := bytes.ToLower(s) //拷贝出小写字母
	fmt.Println(s2, s)

	s2 = bytes.Repeat(s, 2) //返回count个相连字符串
	fmt.Println(s2)

	s = bytes.Replace(s, []byte{65}, []byte{66}, 1) //s是一个字符串 old 是被替换的字符串， new 替换成的字符串 n  替换的次数  小于0全部替换
	fmt.Println(s)

	s = bytes.Trim(s, "B") //去掉给定参数字符
	fmt.Println(s)
	s = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2}
	s5 := bytes.Split(s, []byte{1}) //更具给定的参数切割字符串
	fmt.Println(s5)
	s = bytes.Join(s5, []byte{2}) //以参数连接字符串
	fmt.Println(s)
	r1 := bytes.NewReader([]byte{65, 66, 67})
	r1.Read([]byte{68})
	file, err := os.OpenFile("D:/HxyGo/src/the-way-to-go/bytes/123.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal("文件不存在")
	}
	r1.WriteTo(file) //写入
}
