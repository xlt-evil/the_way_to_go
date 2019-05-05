package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
	c.mux.Unlock()
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
	s := new(SS)
	fmt.Println(s)
	//new 与make的区别 new 只是给了一个指针，并初始化了一下基本数据 而make则是直接返回数据结构 new就好比去前台拿钥匙再去指定地方拿已经初始化好的数据，而make就是到指定的地方拿初始化的东西
	//new(Type)和 make(*Type) 这两个的效果是一样的
	p := new([]int)
	length := len(*p)
	fmt.Println(length)
	fmt.Println(p)
	sss := make([]int, 5, 10)
	pp := sss
	fmt.Println(len(pp))
}

type SS struct {
	money int
}
