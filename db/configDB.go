package db

import (
	"GinWeb/config"
	"GinWeb/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var connect *gorm.DB

func init() {
	db, err := gorm.Open(config.DatabaseConfig.Dialect, config.DatabaseConfig.Database)
	if err != nil {
		panic("failed to connect database")
	}
	if config.ServerConfig.Environment == model.DevelopmentMode {
		db.LogMode(true)
	}
	db.DB().SetMaxIdleConns(config.DatabaseConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.DatabaseConfig.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(config.DatabaseConfig.ConnMaxLifetime))

	// 创建表时添加表后缀
	//db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate((&model.User{}))
	connect = db
}

func Connect() *gorm.DB {
	return connect
}
