package serviceimpl

import (
	"context"
	"fmt"
	"proto-demo/grpc/userservice"
	"time"
)

var UserService = &User{}

type User struct {
	userservice.UnimplementedUserServiceServer
}

func (u *User) GetUser(c context.Context, req *userservice.UserReq) (*userservice.UserResp, error) {
	// 实现具体的业务逻辑
	return &userservice.UserResp{
		Id:   fmt.Sprint(time.Now().UnixMilli()),
		Name: req.Name,
		Age:  req.Age,
	}, nil
}

func (u *User) mustEmbedUnimplementedUserServiceServer() {}
