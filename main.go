package main

import (
	"GinWeb/config"
	"GinWeb/logging"
	"GinWeb/route"
	"github.com/gin-gonic/gin"
)

func main() {
	serverConfig := config.ServerConfig
	logging.InitGinLog()

	router := gin.Default()
	route.Route(router)
	router.Run(serverConfig.Address)
}
