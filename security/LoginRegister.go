package security

import (
	"GinWeb/common"
	"GinWeb/db"
	"GinWeb/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	loginUser := new(model.User)
	if err := c.ShouldBindJSON(loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	LoginCheck(loginUser, c)
	generateToken(loginUser, c)
}

func Register(c *gin.Context) {
	registerUser := new(model.User)
	if err := c.ShouldBindJSON(registerUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := new(model.User)
	connect := db.Connect()
	connect.First(&user, registerUser.LoginId)
	if user.LoginId == registerUser.LoginId && user.Password != registerUser.Password {
		//update password
	} else if user.LoginId == "" {
		//create user
	}else {
		c.JSON(http.StatusOK,common.Error(http.StatusBadRequest,"The user already exists."))
		return
	}
}

func JWTAuth(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		c.JSON(http.StatusUnauthorized, common.AuthError())
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

func LoginCheck(loginUser *model.User, c *gin.Context) {

	user := new(model.User)
	connect := db.Connect()
	connect.First(&user, loginUser.LoginId)

	if user.LoginId == "" {
		c.JSON(http.StatusBadRequest, common.Error(http.StatusBadRequest, "The user does not exist."))
	}
	if loginUser.Password != user.Password || loginUser.LoginId != user.LoginId {
		c.JSON(http.StatusBadRequest, common.Error(http.StatusBadRequest, "Incorrect account or password."))
	}
}
