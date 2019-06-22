package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

func Recover(c *gin.Context) {
	if r := recover(); r != nil {
		if _, ok := r.(runtime.Error); ok {
			panic(r)
		}
		c.JSON(http.StatusOK, Error(http.StatusBadRequest, r.(error).Error()))
	}
}

func PanicError(err string) {
	panic(errors.New(err))
}
