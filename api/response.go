package api

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(c *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	resp := Response{Code: 200, Message: "OK", Data: data}
	c.JSON(http.StatusOK, resp)
}

func HandleError(c *gin.Context, code int, message string, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	resp := Response{Code: code, Message: message, Data: data}
	c.JSON(http.StatusOK, resp)
}
