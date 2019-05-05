package main

import (
	"fmt"
	"time"
)

//学习goroutine 与 channel
//无缓冲的channel 当接受到值时就会把当前的goroutine 挂起来阻塞住了，而有缓冲的则是填满位置才会挂起来
func main() {
	s := make(chan int, 10)
	for i := 10; i > 0; i-- {
		s1 := i //避免发送闭包问题   闭包 1. 是否可以访问其所在作用域的函数 针对的是变量 2.函数嵌套  3.在所在作用域外调用  i明显满足这3个条件因此i变成了类似全局变量的存在
		go func() {
			fmt.Println("这是第", s1, "的goroutine")
			s <- s1

		}()
	}
	time.Sleep(time.Second * 3)
	for i := len(s); i > 0; i-- {
		fmt.Println(<-s)
	}
	//======================================
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c { //会不断的接受channel发送的值，知道被关闭为止，所以如果不关闭的化，并且没有值发送就会发生死锁问题需要注意
		fmt.Println(i)
	}
	//=====================================
	c1 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c1)
		}
		quit <- 0
	}()
	fibonacci_1(c1, quit)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci_1(c, quit chan int) {
	x, y := 0, 1
	for {
		select { //看谁先到谁先匹配
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
