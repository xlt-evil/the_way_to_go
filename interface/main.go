package main

import "fmt"

func main() {
	test_1()
	test_2()
}

//实现一个接口只要重写接口内的方法即可,需要注意的是接受者为指针类型时，原类型时没有实现接口的反之可以
func test_1() {
	var I Myerror
	I.msg = "我的错误是"
	fmt.Println(checkError(I, "错过你"))
}

//接口
type errors interface {
	error(string) string
}

//实现接口的类型
type Myerror struct {
	msg string
}
type Mypoint struct {
	p *int
}

func (this Myerror) error(str string) string {
	return this.msg + str
}

func (this Mypoint) error(str string) string {
	return str
}

func checkError(error errors, str string) string {
	return error.error(str)
}

//类型转换/为nil的类型在var的初始化要加等号
//类型断言的语法是x.(T) T被断言的类型 可以为任意类型 x是接口类型的引用
//类型断言是x的动态值是不是断言类型，所以说x是在之前的时候已经被赋值了，然后进行判断，重点的是接口得动态值类型
//接口的nil 接口分为类型和值 只有类型和值都为nil 时才为nil 所以说当值为nil的时候判断nil 会为false是因为类型不为nil
func test_2() {
	var I = Myerror{"2018 888"}
	var err errors
	var p Mypoint
	err = p
	err = I
	fmt.Println(err == nil)
	fmt.Println(err)
	if s, ok := err.(Myerror); ok { //如果动态值时该类型就发生强转/
		err = I
		result := checkError(err, "开心过大年哈哈哈")
		fmt.Println(result)
		fmt.Println(checkError(s, "6666"))
	} else {
		fmt.Println("BUG")
	}
}
