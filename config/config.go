package config

import (
	"GinWeb/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"unicode/utf8"
)

type databaseConfig struct {
	Dialect         string
	Database        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

var DatabaseConfig databaseConfig

func initDatabaseConfig() {
	util.SetStructByJSON(&DatabaseConfig, configData["database"].(map[string]interface{}))
}

type serverConfig struct {
	Environment string
	LogDir      string
	LogFile     string
	Address     string
}

var ServerConfig serverConfig

func initServerConfig() {
	util.SetStructByJSON(&ServerConfig, configData["server"].(map[string]interface{}))
	sep := string(os.PathSeparator)
	execPath, _ := os.Getwd()
	length := utf8.RuneCountInString(execPath)
	lastChar := execPath[length-1:]
	if lastChar != sep {
		execPath = execPath + sep
	}
	ymdStr := util.GetTodayYMD("-")
	if ServerConfig.LogDir == "" {
		ServerConfig.LogDir = execPath
	} else {
		length := utf8.RuneCountInString(ServerConfig.LogDir)
		lastChar := ServerConfig.LogDir[length-1:]
		if lastChar != sep {
			ServerConfig.LogDir = ServerConfig.LogDir + sep
		}
	}
	ServerConfig.LogFile = ServerConfig.LogDir + ymdStr + ".log"

}

var configData map[string]interface{}

func initConfig() {
	bytes, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println("init config: ", err.Error())
		os.Exit(-1)
	}

	reg := regexp.MustCompile(`/\*.*\*/`)
	configStr := reg.ReplaceAllString(string(bytes[:]), "")
	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &configData); err != nil {
		fmt.Println("invalid config: ", err.Error())
		os.Exit(-1)
	}
}

func initLog() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	logFile, err := os.OpenFile(ServerConfig.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	gin.DefaultWriter = io.MultiWriter(logFile)
}

func init() {
	initConfig()
	initServerConfig()
	initDatabaseConfig()
	initLog()
}
