package main

import (
	"GinWeb/config"
	"GinWeb/route"
	"github.com/gin-gonic/gin"
)

func main() {
	serverConfig :=config.ServerConfig

	router := gin.Default()
	route.Route(router)
	router.Run(serverConfig.Address)
}