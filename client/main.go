package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"proto-demo/client/auth"
	"proto-demo/grpc/userservice"
)

func main() {
	token := &auth.Authentication{
		Token: "123",
	}
	// 1. 新建连接 （没有证书会报错）
	conn, err := grpc.Dial(":8888", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithPerRPCCredentials(token))
	if err != nil {
		log.Fatal(err)
	}
	// 退出时关闭链接
	defer conn.Close()
	// 2. 调用user.pb.go中的NewProdServiceClient方法
	client := userservice.NewUserServiceClient(conn)
	// 3. 直接像调用本地方法一样调用GetUser方法
	resp, err := client.GetUser(context.Background(), &userservice.UserReq{Name: "laoli", Age: 77})
	if err != nil {
		log.Fatal("调用gRPC方法错误: ", err)
	}

	r, _ := json.Marshal(resp)
	fmt.Println("调用gRPC方法成功，response = ", string(r))
	//// 序列化：请求参数转换为二进制
	//bytes, err := proto.Marshal(user)
	//if err == nil {
	//	fmt.Printf("序列化错误：%v", err)
	//	return
	//}
}
