package main

import (
	"IMsystem/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//测试数据库连接

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(0.0.0.0:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"),&gorm.Config{})
	if err!=nil {
		panic("failed to connect database")
	}
	//迁移schema
	db.AutoMigrate(&models.UserBasic{})	

	user := &models.UserBasic{}
	user.Name = "张三"
	db.Create(user)

	//Read
	fmt.Printf("db.First(user, 1): %v\n", db.First(user, 1))

	//修改密码
	db.Model(user).Update("PassWord","1234")

	
}
