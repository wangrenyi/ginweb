package hanlder

import (
	"GinWeb/common"
	"GinWeb/db"
	"GinWeb/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func SaveUser(c *gin.Context) {
	user := new(model.User)
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusOK, common.Error(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}
	user.CreateTime = time.Now()
	user.CreateUser = "admin"
	user.UpdateTime = time.Now()
	user.UpdateUser = "admin"

	connect := db.Connect()
	connect.Create(user)
	c.JSON(http.StatusOK, common.Success("you are logged in"))
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	user := new(model.User)
	connect := db.Connect()
	connect.First(&user, id)
	c.JSON(http.StatusOK, common.Success(user))
}

func GetUsers(c *gin.Context) {
	pageQuery := new(common.PageQuery)
	if err := c.ShouldBindQuery(pageQuery); err != nil {
		c.JSON(http.StatusOK, common.Error(http.StatusBadRequest, err.Error()))
		return
	}
	fmt.Println(pageQuery)

	pageIndex := pageQuery.PageIndex
	pageSize := pageQuery.PageSize
	size := 0
	if pageIndex > 0 {
		size = pageIndex * pageSize
	}

	var users []model.User
	connect := db.Connect()
	connect.Offset(size).Limit(pageSize).Order(pageQuery.OrderBy, true).Find(&users)
	c.JSON(http.StatusOK, common.Success(&users))
}

func UpdateUser(c *gin.Context) {
	user := new(model.User)
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusOK, common.Error(http.StatusBadRequest, err.Error()))
		return
	}
	user.UpdateTime = time.Now()
	user.UpdateUser = "admin"
	preVersion := user.Version
	user.Version = preVersion + 1
	fmt.Println(user)

	connect := db.Connect()
	connect.Model(&user).Where("version = ?", preVersion).Updates(user)
	c.JSON(http.StatusOK, common.Success(&user))
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	user := new(model.User)
	user.Id, _ = strconv.Atoi(id)
	user.Status = 0
	user.UpdateTime = time.Now()
	user.UpdateUser = "admin"

	connect := db.Connect()
	connect.Delete(&user)
	c.JSON(http.StatusOK, common.Info())
}
