package main

import (
	pb "github.com/suyuanhxx/go-kit-demo/proto"
	"google.golang.org/grpc"
	"log"
	"context"
	"io"
	"time"
)

func main() {
	addr := "0.0.0.0:8080"

	// 不使用认证建立连接
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect server error: %v", err)
	}
	defer conn.Close()

	// 创建 gRPC 客户端实例
	grpcClient := pb.NewUserServiceClient(conn)

	// 调用服务端的函数
	req := pb.UserRequest{Id: 2}

	// one
	resp, err := grpcClient.GetUserInfo(context.Background(), &req)
	if err != nil {
		log.Fatalf("recevie resp error: %v", err)
	}
	log.Printf("[RECEIVED RESPONSE]: %v\n", resp) // 输出响应
	log.Printf("one------------------")           // 输出响应

	// two
	stream, err := grpcClient.GetUserInfo2(context.Background(), &req)
	for {
		resp, err := stream.Recv()
		if err == io.EOF { // 服务端数据发送完毕
			break
		}
		if err != nil {
			log.Fatalf("receive error: %v", err)
		}
		log.Printf("[RECEIVED RESPONSE]: %v\n", resp) // 输出响应
	}
	log.Printf("two------------------") // 输出响应

	// three
	// 向服务端发送流数据
	s1, err := grpcClient.GetUserInfo3(context.Background())

	// 模拟的数据库中有 3 条记录，ID 分别为 1 2 3
	for i := 1; i < 4; i++ {
		err := s1.Send(&pb.UserRequest{Id: int32(i)})
		if err != nil {
			log.Fatalf("send error: %v", err)
		}
	}
	// 接收服务端的响应
	resp, err = s1.CloseAndRecv()
	if err != nil {
		log.Fatalf("recevie resp error: %v", err)
	}
	log.Printf("[RECEIVED RESPONSE]: %v\n", resp) // 输出响应
	log.Printf("three------------------")         // 输出响应

	// four
	// 向服务端发送数据流，并处理响应流
	s2, err := grpcClient.GetUserInfo4(context.Background())
	for i := 1; i < 4; i++ {
		s2.Send(&pb.UserRequest{Id: int32(i)})
		time.Sleep(1 * time.Second)
		resp, err := s2.Recv()
		if err != nil {
			log.Fatalf("resp error: %v", err)
		}
		log.Printf("[RECEIVED RESPONSE]: %v\n", resp) // 输出响应
	}
	log.Printf("four------------------") // 输出响应
}
