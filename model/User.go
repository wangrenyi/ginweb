package model

import (
	_ "github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Id         int       `json:"id"`
	LoginId    string    `json:"loginId"`
	Password   string    `json:"password"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"createTime"`
	CreateUser string    `json:"createUser"`
	UpdateTime time.Time `json:"updateTime"`
	UpdateUser string    `json:"updateUser"`
	Status     int       `json:"status"`
	Version    int       `json:"version"`
}

// 设置User的表名为`user`,不设置为users
func (User) TableName() string {
	return "user"
}
