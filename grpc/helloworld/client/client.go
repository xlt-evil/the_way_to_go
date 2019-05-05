package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "the-way-to-go/grpc/helloworld/helloworld"
	"time"
)

func main() {
	//建立一个grpc的安全连接
	conn, err := grpc.Dial(":50001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()
	//建立一个grpc 的简单连接
	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "hxy"})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r.Name)
	}
	return
}
