package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"proto-demo/grpc/userservice"
	"proto-demo/server/serviceimpl"
)

func main() {
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 验证token逻辑
		err = Auth(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}
	rpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor))

	userservice.RegisterUserServiceServer(rpcServer, serviceimpl.UserService)
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("启动监听出错,%v", err)
	}
	startServerErr := rpcServer.Serve(listen)
	if startServerErr != nil {
		log.Fatalf("启动服务出错,%v", startServerErr)
	}
}

// 认证逻辑
func Auth(ctx context.Context) error {
	// 从context中解析出client请求时携带的metadata数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("token miss")
	}
	var token string
	if val, ok := md["token"]; ok {
		token = val[0]
	}
	if token != "123" {
		return status.Errorf(codes.Unauthenticated, "token不合法")
	}
	return nil
}
