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

func Error(statusCode int, msg string) *Result {
	return &Result{
		"msg":        msg,
		"statusCode": statusCode,
	}
}
