package common

import "net/http"

type Result map[string]interface{}

func Success(data interface{}) *Result {
	return &Result{
		"msg":        "Action successfully done.",
		"statusCode": http.StatusOK,
		"data":       &data,
	}
}

func Info() *Result {
	return &Result{
		"msg":        "Action successfully done.",
		"statusCode": http.StatusOK,
	}
}

func Error(statusCode int, msg string) *Result {
	return &Result{
		"msg":        msg,
		"statusCode": statusCode,
	}
}

func AuthError() *Result {
	return &Result{
		"msg":        "Unauthorized",
		"statusCode": http.StatusUnauthorized,
	}
}

func AuthSuccess(authToken string) *Result {
	return &Result{
		"token":      authToken,
		"statusCode": http.StatusOK,
		"msg":        "Authentication Success",
	}
}
