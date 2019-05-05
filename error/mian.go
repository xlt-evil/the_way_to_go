package main

import (
	"errors"
	"fmt"
)

type data struct{ name int }

func (this *data) Error() string { return "" }

func bad() bool {
	return true
}

//自定义错误返回函数
func test() error {
	var p *data
	fmt.Println(p == nil)
	if bad() {
		return p //type *data value nil interface != nil
	}
	return nil
}

//只是返回错误非空
func test1() error {
	var val error = errors.New("XXX")
	return val
}

func main() {
	var e error = test()
	//fmt.Println(e2 == nil)
	//fmt.Println(e == e2)
	//fmt.Println(reflect.TypeOf(e))
	//fmt.Println(reflect.ValueOf(e))
	//fmt.Println()
	//s := new(*int)
	//fmt.Println(s)
	//var s1 *int
	//fmt.Println(s1)
	if e == nil {
		fmt.Println("e is nil")
	} else {
		fmt.Println("e is not nil")
	}
	var e1 error = test1()
	if e1 == nil {
		fmt.Println("e1 is nil")
	} else {
		fmt.Println("e1 is not nil")
		fmt.Println(e1.Error())
	}
}
