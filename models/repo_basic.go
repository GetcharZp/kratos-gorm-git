package models

import "gorm.io/gorm"

type RepoBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"` // 唯一标识
	Path     string `gorm:"column:path;type:varchar(255);" json:"path"`        // 仓库路径
	Name     string `gorm:"column:name;type:varchar(255);" json:"name"`        // Name
	Desc     string `gorm:"column:desc;type:varchar(255);" json:"desc"`        // Desc
	Star     int64  `gorm:"column:star;type:int(11);default:0;" json:"star"`   // Star
	Type     int    `gorm:"column:type;type:tinyint(1);" json:"type"`          // 类型，{1:公库 0:私库}
}

func (table *RepoBasic) TableName() string {
	return "repo_basic"
}
