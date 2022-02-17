package service

import (
	"awesomeProject1/hello/protobuf"
	"context"
)

//type UserService interface {
//	Login(ctx context.Context, req *protobuf.LoginRequest) (*protobuf.LoginResponse, error)
//}

type userService struct {
	*protobuf.UnimplementedUserServiceServer
}

func NewUserService() protobuf.UserServiceServer {
	return &userService{&protobuf.UnimplementedUserServiceServer{}}
}

func (userService *userService) Login(ctx context.Context, req *protobuf.LoginRequest) (*protobuf.LoginResponse, error) {
	if req.Username == "admin" && req.Password == "123456" {
		resp := &protobuf.LoginResponse{
			Code:    10000,
			Message: "登陆成功",
		}
		return resp, nil
	}

	resp := &protobuf.LoginResponse{
		Code:    100002,
		Message: "登陆失败",
	}
	return resp, nil

}
