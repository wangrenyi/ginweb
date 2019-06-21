package main

import (
	"GinWeb/db"
	"GinWeb/hanlder"
	"GinWeb/security"
	"github.com/gin-gonic/gin"
)

func init() {
	db.InitDB()
}

func main() {
	router := gin.Default()

	routerGroup := router.Group("/appbiz")
	{
		routerGroup.POST("/login", security.Login)
		routerGroup.POST("/register", security.Register)
	}
	routerGroup.Use(security.JWTAuth)

	initUserRouter(routerGroup)

	defer db.Connect().Close()

	router.Run("localhost:9080")
}

func initUserRouter(routerGroup *gin.RouterGroup) {
	v1 := routerGroup.Group("/v1")
	{
		v1.POST("/user/save", hanlder.SaveUser)
		v1.GET("/user/:id", hanlder.GetUser)
		v1.GET("/users", hanlder.GetUsers)
		v1.PUT("/user/update", hanlder.UpdateUser)
		v1.DELETE("/user/:id", hanlder.DeleteUser)
	}
}

/**
GIN:{https://www.jianshu.com/p/a3f63b5da74c, https://github.com/gin-gonic/gin/tree/v1.4.0}
GORM:http://gorm.book.jasperxu.com/
JWT:https://github.com/Wangjiaxing123/JwtDemo
*/
