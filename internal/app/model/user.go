package model

import (
	"gorm.io/gorm"
)

// User 模型
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique;not null" json:"username"`
	Email    string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}

// TableName 指定 User 对应的表名
func (User) TableName() string {
	return "user"
}
