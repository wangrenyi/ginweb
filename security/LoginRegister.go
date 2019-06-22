package security

import (
	"GinWeb/common"
	"GinWeb/db"
	"GinWeb/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	defer common.Recover(c)

	loginUser := new(model.User)
	if err := c.ShouldBindJSON(loginUser); err != nil {
		c.JSON(http.StatusOK, common.Error(http.StatusBadRequest, err.Error()))
		return
	}

	LoginCheck(loginUser)

	token, err := generateToken(loginUser)
	if err != nil {
		c.JSON(http.StatusOK, common.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.AuthSuccess(token))
}

func Register(c *gin.Context) {
	defer common.Recover(c)

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
		user.Password = registerUser.Password
		user.Name = registerUser.Name
		connect.Model(user).Updates(user)
	} else if user.LoginId == "" {
		//create user
		registerUser.CreateTime = time.Now()
		registerUser.CreateUser = registerUser.LoginId
		registerUser.UpdateTime = time.Now()
		registerUser.UpdateUser = registerUser.LoginId
		registerUser.Version = 1
		registerUser.Status = 1
		connect.Create(registerUser)
	} else {
		common.PanicError("The user already exists.")
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

func LoginCheck(loginUser *model.User) {

	user := new(model.User)
	connect := db.Connect()
	connect.Where("login_id = ?", loginUser.LoginId).First(&user)

	if user.LoginId == "" {
		common.PanicError("the user does not exist")
	}
	if loginUser.Password != user.Password || loginUser.LoginId != user.LoginId {
		common.PanicError("incorrect account or password")
	}
}
