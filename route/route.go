package route

import (
	"GinWeb/hanlder"
	"GinWeb/security"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {

	routerGroup := router.Group("/appbiz")
	{
		routerGroup.POST("/login", security.Login)
		routerGroup.POST("/register", security.Register)
	}
	routerGroup.Use(security.JWTAuth)

	initUserRouter(routerGroup)
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
