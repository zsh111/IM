package service

import (
	"IMsystem/models"

	"github.com/gin-gonic/gin"
)

//使用gin将信息拿到页面
func GetUserList(c *gin.Context){
	data := make([]*models.UserBasic, 10)
	data =  models.GetUserList()
	c.JSON(200,gin.H{
		"message": data,
	})
}