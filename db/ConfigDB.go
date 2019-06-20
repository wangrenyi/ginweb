package db

import (
	"GinWeb/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var connect *gorm.DB

func InitDB() {
	db, err := gorm.Open("mysql", "root:root@/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(50)
	db.DB().SetConnMaxLifetime(200)

	// 创建表时添加表后缀
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate((&model.User{}))
	db.LogMode(true)
	connect = db
}

func Connect() *gorm.DB {
	return connect
}
