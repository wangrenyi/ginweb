package db

import (
	"GinWeb/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open("mysql", "root:root@/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// 创建表时添加表后缀
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate((&model.User{}))
}

func DBConnect() *gorm.DB {
	return db
}
