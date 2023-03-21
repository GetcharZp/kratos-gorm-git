package service

import (
	"context"
	"fmt"
	"kratos-gorm-git/helper"
	"kratos-gorm-git/models"

	pb "kratos-gorm-git/api/git"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	ub := new(models.UserBasic)
	fmt.Println(req.Username)
	err := models.DB.Where("username = ? AND password = ?", req.Username, helper.GetMd5(req.Password)).First(ub).Error
	if err != nil {
		return nil, err
	}
	token, err := helper.GenerateToken(ub.Identity, ub.Username)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		Token: token,
	}, nil
}
