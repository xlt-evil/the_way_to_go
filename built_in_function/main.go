//内置函数
package main

import (
	"fmt"
)

type Peroson struct {
	s    int
	name string
}

//声明表示在内存中存在了这么一个地址，一般类型在声明的同时就置零了，因为在类型声明同时开辟该类型的内存攻击长度为一内容是为其零值(不给长度的请况下）
//引用类型的声明可以理解为声明了一个指向该类型的指针，声明的同时他们自己初始化，其值为nil表示无，可以理解为出现了一把钥匙没有箱子，所以在他没有指向任何对象前，除了回收操作其余一切都不合法
func main() {
	//New()
	//Make()
	//MyCap()
	//Mycopy()
	//MyDelete()
	//New()
	//Make()
	//MyPrintln()
	//Mypanic()
	//MyClose()
	//Mycomplex()
	//MyReal()
	//Myimag()
}

//把值追加到slice里面，返回修改的slice
//追加的时候，append会先判断增加的slice的容量，超出容量就会把开辟一个新的内存空间用于存放slice然后把追加到该slice后面，如果没有超出，就是追加到slice后面，就是对容量后面的数据进行覆盖
func Myappend() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(cap(num))
	slice := num[:5]
	fmt.Println(cap(slice), len(slice))
	slice = append(slice, 0)
	//因为没有超出范围所以覆盖长度后面的数
	fmt.Println(num)
	slice2 := []int{1, 2, 3, 4, 5, 6, 7}
	slice = append(slice, slice2...)
	//超出容量范围，所以不影响原数组
	fmt.Println(num)
	fmt.Println(slice)
}

//计算类型的上长度度并返回,接受字符串，切片，数组,channel，map 长度指的是已有数据
func MyLen() {
	str := "string"
	fmt.Println(len(str))
	num := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	ss := len(num)
	fmt.Println(ss)
	c := make(chan int, 3)
	//开辟了内存空间，里面没有值
	fmt.Println(len(c))
	c <- 1
	fmt.Println(len(c))
	c <- 2
	fmt.Println(len(c))
	fmt.Println(c)
	return
}

//cap是计算类型的最大长度，接受字符串，切片，数组,channel，长度指的是已有数据
func MyCap() {
	sum := []int{1, 2, 4, 5, 6}
	fmt.Println(cap(sum))
	c := make(chan int, 10)
	fmt.Println(cap(c))
}

//new  与 make的区别
//1. new 是用来分配内存的，他返回的是一个新分配类型的零值的指针 例如 new 一个 int 返回指针取出其类型的零值结果为0 ，而 []int 为slice 零值为nil 去值后为[]==nil true,开辟该类型的内存返回指针
//new 在引用类型 的使用与 var 声明没有区别 只是返回的是一个指针的，一个是对象的
//所以说new是置零，make是初始化
func New() {
	fmt.Println(*new(int))
	fmt.Println(*new([]int))
	fmt.Println(*new([]string))
	fmt.Println(*new([5]int)) //是有初始化内存的
	//chan int  slice map 这3个有
	fmt.Println(*new(chan int))
	var p *[]int = new([]int)
	fmt.Println(*p == nil)
	s := new(map[string]int)
	//s1 := *s
	//s1["name"] = 23
	fmt.Println(*s)
}

//1.只能用来创建slice channel map 并且返回一个初始化的类型为T的值（可以置零） 因为有长度和容量 准备好以后初始化这些值，这些值是零值，并且返回的是一个变量而不是指针
//2，make是在声明引用的同时初始化了引用所指向的对象
func Make() {
	s := make([]int, 0)
	fmt.Println(&s) //当你初始值给零值的时候就相当于置零与new的作用相同
	fmt.Println()
	var ss [5]*int  //声明了指针数组
	var sss *[5]int //声明了数组指针
	fmt.Println(ss)
	fmt.Println(sss)
	//map不能声明容量的,后面
	ssss := make(map[string]int)
	fmt.Println(len(ssss))
	fmt.Println(ssss)
	ssss["213"] = 10
	fmt.Println(ssss)
	//sssss := &ssss["213"]对单独的map值是不能取地址的，因为在扩容时地址的值是会改变的 var s map[string]int = map[string]int{}
}

//把后一个参数赋值给前一个参数返回短的数组长度,接受类型切片，当被赋值的切片小与来源切片时，不会被扩容，多少长度，复制多少
func Mycopy() {
	slice := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s := slice[2:8]
	s1 := slice[3:8]
	fmt.Println(len(s))
	fmt.Println(len(s1))
	copy(s, s1)
	fmt.Println(slice)
	s2 := [10]int{1, 2}
	s3 := [1]int{}
	copy(s3[:], s2[:])
	fmt.Println(s3)
}

//用与删除map中具有指定key的元素
func MyDelete() {
	s := make(map[string]int)
	s["name"] = 1
	v, ok := s["name"] //判断是否有该元素
	fmt.Println(v, ok)
	delete(s, "name")
	v, ok = s["name"]
	fmt.Println(v, ok)
	s["name"] = 2
	fmt.Println(s)
	var s1 Peroson
	fmt.Println(s1.s)
	var m map[string]int //因为map是引用类型所以声明他相当与只声明了一个map型的指针，而引用类型的零值为nil,所以需要给他指向一个内存空间
	fmt.Println(m)
	var slices []int
	fmt.Println(slices == nil)
	//var c chan int
	//fmt.Println(len(c))
	//go func(){
	//	c<-1
	//}()
	//fmt.Println(<-c) //没有指向对象之前的一切操作都不合法
}

//println 与print是输出到标准错误可以理解为他们输出的是不同的页面，但是页面和二为一了导致的结果是与标准输出对的结果混合在一起了
func MyDoublePrintln() {
	str := []string{"name", "hxy"}
	println(str[1], 1)
	print(str[0], 1)
	fmt.Println()
	fmt.Println(&str)
	str[1] = "asd"
}

//该函数会停止当前的goroutine,如果panic不被捕获就会导致当前的goroutine停止，如果捕获则能继续运行
func Mypanic() {
	s := Peroson{}
	defer func() {
		fmt.Println(123)
		if err := recover(); err != nil {
			fmt.Println("出错了", err)
		}
	}()
	panic(s)
}

//捕获的格式 ，捕获到的内容是panic的对象
func Myrecover() {
	defer func() {
		if err := recover(); err != nil {
			//处理的结果
			//现在panic不会向上传播
		}
	}()
}

//用来关闭一个通道，该通道必须是双向的或只能发送的，他只能由发送者执行，而不能有接受者执行，并且关闭之后里面的值依旧可以发送出来。关闭的管道不能重复关闭 。
//总结关闭通道就相当与关闭了接受值的那个接口 以变量的角度就是关闭了发送通道
func MyClose() {
	var s chan<- int //发送 //变量 //接受//对管道来说是接受   //一般在外面说的都是以变量为参照
	//var ss <- chan int//接受 //变量  //发送 //管道
	c := make(chan int, 5)
	s = c
	c <- 1
	c <- 2
	close(s)
	//close(ss)//语法错误
	//close(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(len(c))
	x, ok := <-c
	fmt.Println(x, ok) //可以判断是否通道已关闭
}

//生成一个复数 floatType代表float32/64   complexType 代表 complex64/128   complex64代表float32，complex128代表float64
func Mycomplex() {
	i, j := 1.2, 1.3 //float64
	s := complex(i, j)
	fmt.Println(s)
}

//用与复数中返回的实部
func MyReal() {
	i, j := 1.2, 1.3
	s := complex(i, j)
	l := real(s)
	fmt.Println(l)
}

//用与复数中返回虚部
func Myimag() {
	i, j := 1.2, 1.3
	s := complex(i, j)
	l := imag(s)
	fmt.Println(l)
}

//内置error接口

var err error
