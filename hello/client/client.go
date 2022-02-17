package main

import (
	"awesomeProject1/hello/protobuf"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 2. 创建gRPC客户端
	client := protobuf.NewUserServiceClient(conn)

	// 3. 构造请求参数
	req := &protobuf.LoginRequest{
		Username: "admin",
		Password: "123456",
	}

	// 4. 调用服务端提供的服务
	response, _ := client.Login(context.Background(), req)
	fmt.Println("Login Response: ", response)

}
