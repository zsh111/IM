package models

import (
	"IMsystem/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

//基本信息，用户端
type UserBasic struct {
	gorm.Model
	Name string
	PassWord string
	Phone string
	Email string
	Identity string
	ClientIP string
	ClientPort string
	LoginTime time.Time
	HeartBeatTime time.Time
	LoginOutTime time.Time `gorm:"column:login_out_time" json:"login_out_time"`//修改数据库显示格式
	IsLogout bool
	DeviceInfo string
}
//补充相关方法
func (table *UserBasic) TableName()string{
	return "user_basic"
}

//测试连接数据库,结果存入
func GetUserList()[]*UserBasic{
	data := make([]*UserBasic,10)
	utils.DB.Find(&data)
	for _,v := range data{
		fmt.Println(v)
	}
	return data
}