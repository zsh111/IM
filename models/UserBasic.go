package models

import (
	"IMsystem/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// 基本信息，用户端
type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"` //使用正则匹配电话号码
	Email         string `valid:"email"`
	Identity      string //用于表示token
	ClientIP      string
	ClientPort    string
	Salt          string //加密
	LoginTime     time.Time
	HeartBeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"` //修改数据库显示格式
	IsLogout      bool
	DeviceInfo    string
}

// 补充相关方法
func (table *UserBasic) TableName() string {
	return "user_basic"
}

// 测试连接数据库,结果存入
func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func FindUserByName(name string) UserBasic {
	//按name查找，只返回第一个找到的结果
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func FindUserByEmail(email string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("email = ?", email).First(&user)
	return user
}

func FindUserByPhone(Phone string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("phone = ?", Phone).First(&user)
	return user
}

// 登录逻辑
func FindUserByNameAndPwd(name string, password string) UserBasic {
	//按name查找，只返回第一个找到的结果
	user := UserBasic{}
	utils.DB.Where("name=? and pass_word=?", name, password).First(&user)
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}

/*----------下面实现增删改查----------*/
//增
func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

// 删
func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

// 修改，需要指定修改字段，并未生效
func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Where("id=?", user.ID).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord, Email: user.Email, Phone: user.Phone})
}

// 查询 TODO
func GetUser(name string, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name=? and pass_word=?", name, password).First(&user)
	return user
}
