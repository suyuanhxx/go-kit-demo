package main

import (
	pb "github.com/suyuanhxx/go-kit-demo/proto"
	"context"
	"net"
	"log"
	"google.golang.org/grpc"
	"io"
)

var users = map[int32]pb.UserResponse{
	1: {UserName: "Dennis MacAlistair Ritchie", Age: 70},
	2: {UserName: "Ken Thompson", Age: 75},
	3: {UserName: "Rob Pike", Age: 62},
}

type server struct {
}

func (s *server) GetUserInfo(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{Id: in.Id, UserName: "111"}, nil
}

func (s *server) GetUserInfo2(req *pb.UserRequest, stream pb.UserService_GetUserInfo2Server) error {
	// 响应流数据
	for _, user := range users {
		stream.Send(&user)
	}
	log.Printf("[RECEIVED REQUEST]: %v\n", req)
	return nil
}

func (s *server) GetUserInfo3(stream pb.UserService_GetUserInfo3Server) error {
	var lastID int32
	for {
		req, err := stream.Recv()
		// 客户端数据流发送完毕
		if err == io.EOF {
			// 返回最后一个 ID 的用户信息
			if u, ok := users[lastID]; ok {
				stream.SendAndClose(&u)
				return nil
			}
		}
		lastID = req.Id
		log.Printf("[RECEVIED REQUEST]: %v\n", req)
	}
	return nil
}

func (s *server) GetUserInfo4(stream pb.UserService_GetUserInfo4Server) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		u := users[req.Id]
		err = stream.Send(&u)
		if err != nil {
			return err
		}
		log.Printf("[RECEVIED REQUEST]: %v\n", req)
	}
	return nil
}

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
