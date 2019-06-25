package config

import (
	"GinWeb/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"unicode/utf8"
)

func generateLogFile(logDir string, logFile string) string {
	sep := string(os.PathSeparator)
	execPath, _ := os.Getwd()
	length := utf8.RuneCountInString(execPath)
	lastChar := execPath[length-1:]
	if lastChar != sep {
		execPath = execPath + sep
	}
	ymdStr := util.GetTodayYMD("-")
	if logDir == "" {
		logDir = execPath
	} else {
		length := utf8.RuneCountInString(ServerConfig.LogDir)
		lastChar := logDir[length-1:]
		if lastChar != sep {
			logDir = logDir + sep
		}
	}
	return logDir + ymdStr + "-" + logFile
}

type databaseConfig struct {
	Dialect         string
	Database        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
	LogDir          string
	LogFile         string
}

var DatabaseConfig databaseConfig

func initDatabaseConfig() {
	util.SetStructByJSON(&DatabaseConfig, configData["database"].(map[string]interface{}))
	DatabaseConfig.LogFile = generateLogFile(DatabaseConfig.LogDir, DatabaseConfig.LogFile)
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
	ServerConfig.LogFile = generateLogFile(ServerConfig.LogDir, ServerConfig.LogFile)
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

func init() {
	initConfig()
	initServerConfig()
	initDatabaseConfig()
}
