package logging

import (
	"GinWeb/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"io"
	"log"
	"os"
	"time"
)

func InitDBLog(db *gorm.DB) {
	db.LogMode(true)
	logFile, err := os.OpenFile(config.DatabaseConfig.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	prefix := "[Gorm] " + time.Now().Format("2006-01-02 15:04:05") + " "
	db.SetLogger(log.New(logFile, prefix, 0))
}

func InitGinLog() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	logFile, err := os.OpenFile(config.ServerConfig.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout) //日志文件和控制台上显示
}
