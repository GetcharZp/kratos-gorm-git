package service

import (
	"context"
	"errors"
	pb "kratos-gorm-git/api/git"
	"kratos-gorm-git/helper"
	"kratos-gorm-git/models"
)

type RepoService struct {
	pb.UnimplementedRepoServer
}

func NewRepoService() *RepoService {
	return &RepoService{}
}

func (s *RepoService) CreateRepo(ctx context.Context, req *pb.CreateRepoRequest) (*pb.CreateRepoReply, error) {
	// 1. 查重
	var cnt int64
	err := models.DB.Model(new(models.RepoBasic)).Where("path = ?", req.Path).Count(&cnt).Error
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("路径已存在")
	}
	// 2. 落库
	rb := &models.RepoBasic{
		Identity: helper.GetUUID(),
		Path:     req.Path,
		Name:     req.Name,
		Desc:     req.Desc,
		Type:     int(req.Type),
	}
	err = models.DB.Create(rb).Error
	if err != nil {
		return nil, err
	}
	return &pb.CreateRepoReply{}, nil
}
func (s *RepoService) UpdateRepo(ctx context.Context, req *pb.UpdateRepoRequest) (*pb.UpdateRepoReply, error) {
	return &pb.UpdateRepoReply{}, nil
}
func (s *RepoService) DeleteRepo(ctx context.Context, req *pb.DeleteRepoRequest) (*pb.DeleteRepoReply, error) {
	return &pb.DeleteRepoReply{}, nil
}
func (s *RepoService) GetRepo(ctx context.Context, req *pb.GetRepoRequest) (*pb.GetRepoReply, error) {
	return &pb.GetRepoReply{}, nil
}
func (s *RepoService) ListRepo(ctx context.Context, req *pb.ListRepoRequest) (*pb.ListRepoReply, error) {
	rb := make([]*models.RepoBasic, 0)
	var cnt int64
	err := models.DB.Model(new(models.RepoBasic)).Count(&cnt).Offset(int((req.Page - 1) * req.Size)).Limit(int(req.Size)).
		Find(&rb).Error
	if err != nil {
		return nil, err
	}
	list := make([]*pb.ListRepoItem, 0, len(rb))
	for _, v := range rb {
		list = append(list, &pb.ListRepoItem{
			Identity: v.Identity,
			Name:     v.Name,
			Desc:     v.Desc,
			Path:     v.Path,
			Star:     v.Star,
		})
	}
	return &pb.ListRepoReply{
		List: list,
		Cnt:  cnt,
	}, nil
}
