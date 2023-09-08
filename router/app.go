package router

import (
	"IMsystem/service"

	"github.com/gin-gonic/gin"
)

//用作浏览器访问
func Router() *gin.Engine{
	r := gin.Default()
	r.GET("/index",service.GetIndex)
	r.GET("/user/getUserList",service.GetUserList)
	return r
}