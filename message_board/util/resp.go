package util

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模板
type respTemplate struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

var OK = respTemplate{
	Status: 200, //成功
	Info:   "success",
}

var ParamError = respTemplate{
	Status: 300, //重定向
	Info:   "params error",
}

var RequestError = respTemplate{
	Status: 400, //请求错误
	Info:   "request error",
}

var InternalError = respTemplate{
	Status: 500, //服务器错误
	Info:   "internal error",
}

func RespOK(c *gin.Context) {
	c.JSON(http.StatusOK, OK) //200
}

func ResParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ParamError) //返回参数错误,服务器不理解请求 400
}
func Err(c *gin.Context) {
	c.JSON(http.StatusLengthRequired, RequestError) //需要有效长度
}

func RespInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, InternalError) //服务器内部错误，遇到未曾预料的情况500
}

func ResInterErr(c *gin.Context) {
	c.JSON(http.StatusAlreadyReported, InternalError) //服务器错误
}

// 标准错误
func NormErr(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
}
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}
