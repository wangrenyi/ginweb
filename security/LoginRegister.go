package security

import (
	"GinWeb/common"
	"GinWeb/db"
	"GinWeb/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	loginUser := new(model.User)
	if err := c.ShouldBindJSON(loginUser); err != nil {
		c.JSON(http.StatusOK, common.Error(http.StatusBadRequest, err.Error()))
		return
	}

	if err := LoginCheck(loginUser); err != nil {
		c.JSON(http.StatusOK, common.Error(http.StatusBadRequest, err.Error()))
		return
	}

	token, err := generateToken(loginUser)
	if err != nil {
		c.JSON(http.StatusOK, common.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.AuthSuccess(token))
}

func Register(c *gin.Context) {
	registerUser := new(model.User)
	if err := c.ShouldBindJSON(registerUser); err != nil {
		c.JSON(http.StatusOK, common.Error(http.StatusBadRequest, err.Error()))
		return
	}

	user := new(model.User)
	connect := db.Connect()
	connect.Where("login_id = ?", registerUser.LoginId).First(&user)
	if user.LoginId == registerUser.LoginId && user.Password != registerUser.Password {
		//update password and name
		user.Password=registerUser.Password
		user.Name=registerUser.Name
		connect.Model(user).Updates(user)
	} else if user.LoginId == "" {
		//create user
		registerUser.CreateTime=time.Now()
		registerUser.CreateUser=registerUser.LoginId
		registerUser.UpdateTime=time.Now()
		registerUser.UpdateUser=registerUser.LoginId
		registerUser.Version=1
		registerUser.Status=1
		connect.Create(registerUser)
	} else {
		c.JSON(http.StatusOK, common.Error(http.StatusBadRequest, "The user already exists."))
		return
	}
	c.JSON(http.StatusOK, common.Info())
}

func JWTAuth(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		c.JSON(http.StatusOK, common.AuthError())
		c.Abort()
		return
	}

	j := NewJWT()
	claims, err := j.ParseToken(authToken)

	if err != nil {
		if err == TokenExpired {
			c.JSON(http.StatusOK, common.Error(http.StatusUnauthorized, "Authorization expired"))
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, common.Error(http.StatusUnauthorized, err.Error()))
		c.Abort()
		return
	}

	c.Set("claims", claims)
}

func LoginCheck(loginUser *model.User) error {

	user := new(model.User)
	connect := db.Connect()
	connect.Where("login_id = ?", loginUser.LoginId).First(&user)

	if user.LoginId == "" {
		return errors.New("the user does not exist")
	}
	if loginUser.Password != user.Password || loginUser.LoginId != user.LoginId {
		return errors.New("incorrect account or password")
	}

	return nil
}
