package models

import "gorm.io/gorm"

type RepoUser struct {
	gorm.Model
	Rid  uint `gorm:"column:rid;type:int(11);" json:"rid"`      // 仓库ID
	Uid  uint `gorm:"column:uid;type:int(11);" json:"uid"`      // 用户ID
	Type int  `gorm:"column:type;type:tinyint(1);" json:"type"` // 类型，{1:所有者 2:被授权者}
}

func (table *RepoUser) TableName() string {
	return "repo_user"
}
