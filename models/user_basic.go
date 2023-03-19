package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"`  // 唯一标识
	Username string `gorm:"column:username;type:varchar(255);" json:"username"` // 用户名
	Password string `gorm:"column:password;type:varchar(36);" json:"password"`  // 密码
	Email    string `gorm:"column:email;type:varchar(100);" json:"email"`       // 电子邮箱
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
