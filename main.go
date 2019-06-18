package main

import (
	"GinWeb/db"
	"GinWeb/hanlder"
	"github.com/gin-gonic/gin"
)

func main() {
	initRouter()
}

func init() {
	db.InitDB()
}

func initRouter() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/user/save", hanlder.SaveUser)
		v1.GET("/user/:id", hanlder.GetUser)
		v1.POST("/users", hanlder.GetUsers)
		v1.PUT("/user/update", hanlder.UpdateUser)
		v1.DELETE("/user/:id", hanlder.DeleteUser)
	}
	router.Run()
}
