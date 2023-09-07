package models

import "gorm.io/gorm"

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
	LoginTime uint64
	HeartBeatTime uint64
	LogOutTime uint64
	IsLogout bool
	DeviceInfo string
}
//补充相关方法
func (table *UserBasic) TableName()string{
	return "user_basic"
}