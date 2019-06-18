package hanlder

import (
	"GinWeb/common"
	"GinWeb/db"
	"GinWeb/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SaveUser(c *gin.Context) {
	user := new(model.User)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.CreateTime = time.Now()
	user.CreateUser = "admin"
	user.UpdateTime = time.Now()
	user.UpdateUser = "admin"

	connect := db.DBConnect()
	connect.Create(&user)
	c.JSON(http.StatusOK, common.Success("you are logged in"))
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	user := new(model.User)
	connect := db.DBConnect()
	connect.First(&user, id)
	c.JSON(http.StatusOK, common.Success(&user))
}

func GetUsers(c *gin.Context) {
	//pageIndex := c.DefaultQuery("pageIndex", 0)
	pageQuery := new(common.PageQuery)
	if err := c.ShouldBindJSON(&pageQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pageIndex := pageQuery.PageIndex
	pageSize := pageQuery.PageSize
	size := 0
	if pageIndex > 0 {
		size = int(pageIndex * pageSize)
	}

	var users []model.User
	connect := db.DBConnect()
	connect.Offset(size).Limit(pageSize).Order(pageQuery.OrderBy, true).Find(&users)
	c.JSON(http.StatusOK, common.Success(&users))
}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
