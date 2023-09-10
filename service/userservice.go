package service

import (
	"IMsystem/models"
	"fmt"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @summary 用户列表
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	//使用gin将信息拿到页面，注意注释必须紧贴函数
	data := models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "两次密码不一样!",
		})
		return
	}
	user.PassWord = password
	user.LoginTime = time.Now()
	user.HeartBeatTime = time.Now()
	user.LoginOutTime = time.Now()
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "新增用户成功",
	})
}

// DeleteUser
// @summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "删除用户成功",
	})
}

// UpdateUser
// @summary 修改用户信息
// @Tags 用户模块
// @param id query string false "id"
// @param name query string false "用户名"
// @param password query string false "密码"
// @param email query string false "邮箱"
// @param phone query string false "电话"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	user.Name = c.Query("name")
	user.PassWord = c.Query("password")
	user.Email = c.Query("email")
	user.Phone = c.Query("phone")
	user.LoginTime = time.Now()
	user.HeartBeatTime = time.Now()
	user.LoginOutTime = time.Now()
	_, err := govalidator.ValidateStruct(user) //增加校验规则
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "更新用户失败",
		})
		return
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"message": "更新用户成功",
		})
	}
}
