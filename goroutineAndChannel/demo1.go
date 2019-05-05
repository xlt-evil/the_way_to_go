package main

import (
	//"fmt"
	"sync"
	"time"
	//"strconv"
	"fmt"
)

//创建10个协程 然后顺序输出

var sum int
var l sync.Mutex

func main() {
	c := make(chan int, 1) //
	for i := 0; i < 10; i++ {
		go func() {
			c <- i
			l.Lock()
			fmt.Println(<-c)
			l.Unlock()
		}()
	}
	time.Sleep(3000)

}
