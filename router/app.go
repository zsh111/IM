package router

import (
	"IMsystem/docs"
	"IMsystem/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//用作浏览器访问
func Router() *gin.Engine{
	//使用swag优化前端页面显示，不用init直接访问
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	//下面设置了页面的访问子页面，下面请求与swag的router一致
	r.GET("/swagger/*ang",ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index",service.GetIndex)
	r.GET("/user/getUserList",service.GetUserList)
	r.GET("/user/createUser",service.CreateUser)
	r.GET("/user/deleteUser",service.DeleteUser)
	r.POST("/user/updateUser",service.UpdateUser)//修改请求方式
	return r
}