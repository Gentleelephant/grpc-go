package main

import (
	"awesomeProject1/hello/protobuf"
	"awesomeProject1/hello/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	Address = "127.0.0.1:8888"
)

func main() {
	// 1. 监听端口
	l, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("listen on %s", Address)

	// 2. 实例化grpc服务端
	grpcServcer := grpc.NewServer()

	// 注册实现的服务实例
	userService := service.NewUserService()
	protobuf.RegisterUserServiceServer(grpcServcer, userService)

	log.Printf("gRPC is runnning...")
	err = grpcServcer.Serve(l)
	if err != nil {
		log.Printf("gRPC server err:%s\n", err)
	}

}
