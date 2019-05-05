package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	client, err := rpc.Dial("tcp", "192.168.3.93:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	fmt.Println("rpc调用 ")
	fmt.Println("请输入两个数字数字间以空格为间隔")
	args := &Args{7, 8}
	for {
		fmt.Scanf("%d%d", &args.A, &args.B)
		fmt.Println("远程调用中~")
		var reply int
		err = client.Call("Arith.Multiply", args, &reply)

		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("乘法为: %d*%d=%d\n", args.A, args.B, reply)

		// Asynchronous call
		quotient := new(Quotient)
		divCall := client.Go("Arith.Divide", args, quotient, nil)
		replyCall := <-divCall.Done // will be equal to divCall
		if replyCall.Error != nil {
			log.Fatal("arith error:", replyCall.Error)
		}
		fmt.Printf("除法为: %d/%d=%d...%d", args.A, args.B, quotient.Quo, quotient.Rem)
		// check errors, print, etc.
		fmt.Println()
		fmt.Println("调用结束")
		fmt.Println("输入e结束,其他继续")
		str := ""
		fmt.Scanf("%s", &str)
		fmt.Println(str)
		if str == "e" {
			break
		}
	}
}
