package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	pb "the-way-to-go/grpc/helloworld/helloworld"
)

const port = ":50001"

//server 要实现helloworld.greeterserver 的接口
type server struct {
}

//实现接口
func (t *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Name: in.Name + "welcome to grpc world "}, nil
}

//rpc 是在运输层的协议，rpc 假定 传输协议已存在（tcp,udp）
// 远程服务调用
// rpc 一般运用与内部传输
//开启服务
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("服务开启失败具体原因:", err)
		return
	}
	// 开启
	s := grpc.NewServer()
	//注册上一个s
	pb.RegisterGreeterServer(s, &server{})
	//将服务反射到grpc服务上，理解为注册路由 = =
	reflection.Register(s)
	//监听服务
	if err := s.Serve(lis); err != nil {
		fmt.Println("监听失败", err.Error())
	}
}
