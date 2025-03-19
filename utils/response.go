package utils

import "github.com/gin-gonic/gin"

// Response 通用响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// SuccessResponse 返回成功响应
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, Response{Code: 200, Msg: "成功", Data: data})
}

// ErrorResponse 返回错误响应
func ErrorResponse(c *gin.Context, msg string) {
	c.JSON(400, Response{Code: 400, Msg: msg})
}
