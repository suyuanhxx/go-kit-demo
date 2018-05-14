package main

import (
	pb "github.com/suyuanhxx/go-kit-demo/proto"
	"net"
	"log"
	"google.golang.org/grpc"
)

func main() {

	addr := "0.0.0.0:8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	} else {
		log.Println("server listen: ", addr)
	}

	// 创建 gRPC 服务器实例
	grpcServer := grpc.NewServer()

	// 向 gRPC 服务器注册服务
	pb.RegisterUserServiceServer(grpcServer, &server{})

	// 启动 gRPC 服务器
	// 阻塞等待客户端的调用
	grpcServer.Serve(listener)
}
