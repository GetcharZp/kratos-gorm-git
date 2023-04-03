package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	pb "kratos-gorm-git/api/git"
	"kratos-gorm-git/define"
	"kratos-gorm-git/helper"
	"kratos-gorm-git/models"
	"os"
	"os/exec"
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
	err = models.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(rb).Error
		if err != nil {
			return err
		}
		// init repo path
		// mkdir path
		gitRepoPath := define.RepoPath + string(os.PathSeparator) + req.Path
		err = os.MkdirAll(gitRepoPath, 0755)
		if err != nil {
			return err
		}
		// git init --bare
		cmd := exec.Command("/bin/bash", "-c", "cd "+gitRepoPath+" ; git init --bare")
		err = cmd.Run()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateRepoReply{}, nil
}
func (s *RepoService) UpdateRepo(ctx context.Context, req *pb.UpdateRepoRequest) (*pb.UpdateRepoReply, error) {
	err := models.DB.Model(new(models.RepoBasic)).Where("identity = ?", req.Identity).Updates(map[string]interface{}{
		"name": req.Name,
		"desc": req.Desc,
		"type": req.Type,
	}).Error
	if err != nil {
		return nil, err
	}
	return &pb.UpdateRepoReply{}, nil
}
func (s *RepoService) DeleteRepo(ctx context.Context, req *pb.DeleteRepoRequest) (*pb.DeleteRepoReply, error) {
	// 1. 获取仓库的基础信息
	var rb = new(models.RepoBasic)
	err := models.DB.Model(new(models.RepoBasic)).Where("identity = ?", req.Identity).First(rb).Error
	if err != nil {
		return nil, err
	}
	// 2. 删除记录
	err = models.DB.Transaction(func(tx *gorm.DB) error {
		// 2.1 删除仓库数据
		err = os.RemoveAll(define.RepoPath + string(os.PathSeparator) + rb.Path)
		if err != nil {
			return err
		}
		// 2.2 删除DB记录
		err = tx.Where("identity = ?", req.Identity).Delete(new(models.RepoBasic)).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
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
