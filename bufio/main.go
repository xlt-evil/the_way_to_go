// 学习bufio包 该包实现了
//1.有缓冲的I/O 包装了一个io.Reader 和io.Writer 接口对象
//2.提供一些缓冲和一些文本I/O的帮助函数对象
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	MyScanner()
	s := strings.NewReader("ABC\nDEF\r\nGHI\nJKL")
	bs := bufio.NewScanner(s)
	for bs.Scan() {
		fmt.Printf("%s %v\n", bs.Bytes(), bs.Text())
	}
}

//读
func MyReader() {
	r := bufio.Reader{} //7个变量//该对象实现了给io.Reader接口对象附加缓冲
	file, _ := os.OpenFile("D:/HxyGo/src/the-way-to-go/bufio/123.txt", os.O_WRONLY, os.ModePerm)
	r1 := bufio.NewReader(file) //创建一个具有默认缓冲大小，从r出读取的reader

	r2 := bufio.NewReaderSize(file, 100) //不能有两个流同时对内存中的数据进行读取？ //创建一个给定缓存大小的reader

	b, _, _ := r1.ReadLine() ////readline是一个低水平的行读取原语，建议使用readByte("\n"),readstring("\n")或者scanner。该方法超出缓存，返回值2位true第二次调用继续读取该行
	fmt.Println(string(b))

	left := r1.Buffered() //缓存区可读字节数 //就是流没有读完的字节数
	fmt.Println(left)     //剩余未读取字节

	n, _ := r1.Peek(1) //返回输入流的下n个字节，但不会移动读取位置
	fmt.Println(string(n))

	b1 := make([]byte, 10)
	r1.Read(b1) //读取给定的字节量
	fmt.Println(string(b1))

	bs, _ := r1.ReadByte() //读取一个字节
	fmt.Println(string(bs))

	ru, n1, _ := r1.ReadRune() //读取一个utf-8字节第一个数码点，第二个是字节量因为utf-8是可变字节
	fmt.Println(ru, n1)

	//r1.UnreadByte()//UnreadRune//撤销最后读取的一个字节/utf-8

	b, _ = r1.ReadSlice(57) //读取到给定的码点时停止，并返回+停止的码点，在没读到码点之前缓冲填满会返回一个错误ErrBufferFull,下一次I/O操作重写，所以尽量使用 rBytes/rString 读取
	fmt.Println(string(b))

	r1.Reset(file) //丢弃缓冲区的数据，清除错误重置,把流的缓存数据清空了 可以切换读取对象
	fmt.Println(r)
	fmt.Println(r2)

	var rus rune
	rus = 57
	fmt.Println("rune", string(rus))
	//r1.ReadBytes/ReadString(delim)//读取到给定码点返回+码点的切片，但没有读到码点会返回读取到的数据和错误，只有读取不到delim时返回才会返回一个nil错误 ,这两个函数只是返回类型不同而已

}

//写
func MyWriter() {
	//write
	file, _ := os.OpenFile("D:/HxyGo/src/the-way-to-go/bufio/123.txt", os.O_WRONLY, os.ModePerm)
	file2, err := os.OpenFile("D:/HxyGo/src/the-way-to-go/bufio/453.txt", os.O_WRONLY, os.ModePerm) //给权限才能写否则会Access is denied
	fmt.Println(err)
	r3 := bufio.NewWriter(file2) //创建有两个函数与reader没区别

	b := []byte("我是大神")
	nn, err := r3.Write(b) //Write将b的内容写入 缓冲 。返回写入的字节数。如果返回值nn < len(p)(就说么没有全部写入)，还会返回一个错误说明原因。
	fmt.Println(nn)
	fmt.Println()
	r3.WriteString("气温较为巨额偶奇") //与write没区别只是接受的参数不同
	err = r3.Flush()           //把缓存的数据输出到流
	fmt.Println(err)
	r3.Reset(file) //丢弃缓存中的数据，
	r3.WriteString("hahahaha")
	err = r3.Flush()
	r3.Available() //还有多少字节没写入
	r3.Buffered()  //已写入多少字节
	//其他方法基本没区别
	fmt.Println(err)
}

//读写
func MyRWer() {
	//便于管理就是集成了集合了钱两个类型的方法
	file, _ := os.OpenFile("D:/HxyGo/src/the-way-to-go/bufio/123.txt", os.O_WRONLY, os.ModePerm)
	file2, _ := os.OpenFile("D:/HxyGo/src/the-way-to-go/bufio/453.txt", os.O_WRONLY, os.ModePerm) //给权限才能写否则会Access is denied
	r1 := bufio.NewReader(file)                                                                   //创建一个具有默认缓冲大小，从r出读取的reader
	r3 := bufio.NewWriter(file2)                                                                  //创建有两个函数与reader没区别
	//创建一个reader/writer
	r4 := bufio.NewReadWriter(r1, r3)
	fmt.Println(r4)
}

//分割函数，用于分割文本读取的
func MySplitFunc() {

	// SplitFunc 用来定义“切分函数”类型
	// data 是要扫描的数据
	// atEOF 标记底层 io.Reader 中的数据是否已经读完
	// advance 返回 data 中已处理的数据长度
	// token 返回找到的“指定部分”
	// err 返回错误信息
	//如果在 data 中无法找到一个完整的“指定部分”
	// 则 SplitFunc 返回 (0, nil) 来告诉 Scanner
	// 向缓存中填充更多数据，然后再次扫描
	//
	// 如果返回的 err 是非 nil 值，扫描将被终止，并返回错误信息
	//
	// 如果 data 为空，则“切分函数”将不被调用
	// 意思是在 SplitFunc 中不必考虑 data 为空的情况
}

//Scanner类型提供了方便的读取数据的接口，如从换行符分隔的文本里读取每一行。
func MyScanner() {
	file, _ := os.Open("D:/HxyGo/src/the-way-to-go/bufio/123.txt")
	bs := make([]byte, 100)
	file.Read(bs)
	fmt.Println(bs)
	r1 := bufio.NewScanner(file)
	//r1.Split(bufio.ScanRunes)//设置分割函数必须在scan之前设置
	b := r1.Scan()
	err := r1.Err()
	fmt.Println(err)
	fmt.Println(b)
	for r1.Scan() { //读取到token 并移动到下一个token没有了就false
		fmt.Println(r1.Text(), r1.Bytes()) //都是当前的token 会随着scan的调用而改变
	}
	fmt.Println(r1.Bytes(), r1.Text())
}
