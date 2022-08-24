package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"proto-demo/grpc/userservice"
	"proto-demo/server/serviceimpl"
)

func main() {
	rpcServer := grpc.NewServer()
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
