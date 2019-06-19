package main

import (
	"GinWeb/db"
	"GinWeb/hanlder"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	initUserRouter(router)

	router.Run("localhost:9080")
}

func init() {
	db.InitDB()
}

func initUserRouter(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	{
		v1.POST("/user/save", hanlder.SaveUser)
		v1.GET("/user/:id", hanlder.GetUser)
		v1.GET("/users", hanlder.GetUsers)
		v1.PUT("/user/update", hanlder.UpdateUser)
		v1.DELETE("/user/:id", hanlder.DeleteUser)
	}
}
